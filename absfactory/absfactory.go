package main

import "fmt"

type Shirter interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) getSize() int {
	return s.size
}

type adidasShirt struct {
	shirt
}

type nikeShirt struct {
	shirt
}

type Shoter interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shot struct {
	logo string
	size int
}

func (s *shot) setLogo(logo string) {
	s.logo = logo
}

func (s *shot) setSize(size int) {
	s.size = size
}

func (s *shot) getLogo() string {
	return s.logo
}

func (s *shot) getSize() int {
	return s.size
}

type adidasShot struct {
	shot
}

type nikeShot struct {
	shot
}

type nike struct{}

func (n *nike) makeShot() Shoter {
	return &nikeShot{
		shot: shot{
			logo: "nike",
			size: 12,
		},
	}
}

func (n *nike) makeShirt() Shirter {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 12,
		},
	}
}

type adidas struct{}

func (a *adidas) makeShot() Shoter {
	return &adidasShot{
		shot: shot{
			logo: "adidas",
			size: 111,
		},
	}
}

func (a *adidas) makeShirt() Shirter {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 111,
		},
	}
}

type SportFactory interface {
	makeShot() Shoter
	makeShirt() Shirter
}

func getSportFactory(brand string) (SportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("nothing same brand")
}

func printShirt(s Shirter) {
	fmt.Printf("shirt.logo: %v, shirt.size: %v\n", s.getLogo(), s.getSize())
}

func printShot(s Shoter) {
	fmt.Printf("shot.logo: %v, shot.size: %v\n", s.getLogo(), s.getSize())
}

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
