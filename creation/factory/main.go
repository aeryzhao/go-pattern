package main

import "fmt"

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
