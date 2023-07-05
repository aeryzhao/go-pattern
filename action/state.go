package main

import (
	"fmt"
	"log"
)

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState

	return v
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount += count
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (s *hasItemState) addItem(count int) error {
	s.vendingMachine.incrementItemCount(count)
	return nil
}

func (s *hasItemState) requestItem() error {
	fmt.Println("Item requested")
	s.vendingMachine.setState(s.vendingMachine.itemRequested)
	return nil
}

func (s *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (s *hasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (s *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("item out of stock")
}

func (s *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (s *itemRequestedState) insertMoney(money int) error {
	if money < s.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", s.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	s.vendingMachine.setState(s.vendingMachine.hasMoney)
	return nil
}

func (s *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (s *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (s *hasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (s *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (s *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	s.vendingMachine.itemCount--
	if s.vendingMachine.itemCount == 0 {
		s.vendingMachine.setState(s.vendingMachine.noItem)
	} else {
		s.vendingMachine.setState(s.vendingMachine.hasItem)
	}
	return nil
}

type noItemState struct {
	vendingMachine *vendingMachine
}

func (s *noItemState) addItem(count int) error {
	s.vendingMachine.incrementItemCount(count)
	s.vendingMachine.setState(s.vendingMachine.hasItem)
	return nil
}

func (s *noItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (s *noItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (s *noItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
