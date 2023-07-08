package main

import "fmt"

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restore(m *memento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}

type careTaker struct {
	mementoList []*memento
}

func (c *careTaker) add(m *memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) get(index int) *memento {
	return c.mementoList[index]
}

func main() {
	caretaker := &careTaker{
		mementoList: make([]*memento, 0),
	}

	originator := &originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.add(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.add(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.add(originator.createMemento())

	originator.restore(caretaker.get(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restore(caretaker.get(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
