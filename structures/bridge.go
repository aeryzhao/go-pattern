package main

import "fmt"

type Computer interface {
	Print()
	SetPrinter(printer Printer)
}

type Printer interface {
	PrintFile()
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type HP struct{}

func (p *HP) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {
	hp := &HP{}
	epson := &Epson{}

	mac := &Mac{}
	mac.SetPrinter(hp)
	mac.Print()

	mac.SetPrinter(epson)
	mac.Print()

	win := &Windows{}
	win.SetPrinter(hp)
	win.Print()

	win.SetPrinter(epson)
	win.Print()
}
