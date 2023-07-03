package main

type command interface {
	execute()
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type OnCommand struct {
	device device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type device interface {
	on()
	off()
}

type tv struct{}

func (t *tv) on() {
	println("tv is on")
}

func (t *tv) off() {
	println("tv is off")
}

type ac struct{}

func (a *ac) on() {
	println("ac is on")
}

func (a *ac) off() {
	println("ac is off")
}

func main() {
	tv := &tv{}
	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
	ac := &ac{}
	onCommand = &OnCommand{
		device: ac,
	}
	offCommand = &OffCommand{
		device: ac,
	}
	onButton = &button{
		command: onCommand,
	}
	onButton.press()
	offButton = &button{
		command: offCommand,
	}
	offButton.press()
}
