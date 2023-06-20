package main

type Computer interface {
	insertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) insertIntoLightningPort() {
	println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) insertIntoLightningPort() {
	println("Adapter converts lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}

type client struct{}

func (c *client) insertLightningConnectorIntoComputer(com Computer) {
	println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

func main() {
	client := &client{}
	mac := &Mac{}
	client.insertLightningConnectorIntoComputer(mac)
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}
	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
