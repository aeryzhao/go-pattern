package main

import "fmt"

func main() {
	adidasFactory, _ := getSportFactory("adidas")
	adidasShirt := adidasFactory.makeShirt()
	adidasShot := adidasFactory.makeShot()
	printShirt(adidasShirt)
	printShot(adidasShot)

	nikeFactory, _ := getSportFactory("nike")
	nikeShirt := nikeFactory.makeShirt()
	nikeShot := nikeFactory.makeShot()
	printShirt(nikeShirt)
	printShot(nikeShot)
}

func printShirt(s iShirt) {
	fmt.Printf("shirt.logo: %v, shirt.size: %v\n", s.getLogo(), s.getSize())
}

func printShot(s iShot) {
	fmt.Printf("shot.logo: %v, shot.size: %v\n", s.getLogo(), s.getSize())
}
