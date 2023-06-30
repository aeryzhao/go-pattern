package main

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Department interface {
	Execute(*Patient)
	setNext(department Department)
}

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	println("Reception registering patient")
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.medicineDone {
		println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.Execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		println("Payment done")
		return
	}
	println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

func main() {
	cashier := &Cashier{}
	medical := &Medical{}
	medical.setNext(cashier)
	doctor := &Doctor{}
	doctor.setNext(medical)
	reception := &Reception{}
	reception.setNext(doctor)
	patient := &Patient{name: "abc"}
	reception.Execute(patient)
}
