package main

import "fmt"

type Pizza interface {
	getPrice() int
}

type VeggeMania struct{}

func (p *VeggeMania) getPrice() int {
	return 15
}

type TomotaTopping struct {
	Pizza Pizza
}

func (p *TomotaTopping) getPrice() int {
	return p.Pizza.getPrice() + 3
}

type CheeseTopping struct {
	Pizza Pizza
}

func (p *CheeseTopping) getPrice() int {
	return p.Pizza.getPrice() + 8
}

func main() {
	pizza := &VeggeMania{}
	tomotaTopping := &TomotaTopping{
		Pizza: pizza,
	}
	pizzaWithCheeseAndTomato := &CheeseTopping{
		Pizza: tomotaTopping,
	}
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
