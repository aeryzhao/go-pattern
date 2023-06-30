package main

import "fmt"

type computer interface {
	Print()
	SetPrinter(printer Printer)
}

type Printer interface {
	PrintFile()
}

type mac struct {
	printer Printer
}

func (m *mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *mac) SetPrinter(p Printer) {
	m.printer = p
}

type windows struct {
	printer Printer
}

func (w *windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *windows) SetPrinter(p Printer) {
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

	mac := &mac{}
	mac.SetPrinter(hp)
	mac.Print()

	mac.SetPrinter(epson)
	mac.Print()

	win := &windows{}
	win.SetPrinter(hp)
	win.Print()

	win.SetPrinter(epson)
	win.Print()
}
