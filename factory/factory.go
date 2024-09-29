package main

import (
	"fmt"
)

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "ak47 gun",
			power: 4,
		},
	}
}

type musket struct {
	gun
}

func newmusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newmusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printGun(ak47)
	printGun(musket)
}

func printGun(i iGun) {
	fmt.Printf("Gun: %v\n", i.getName())
	fmt.Printf("Power: %v\n", i.getPower())
}
