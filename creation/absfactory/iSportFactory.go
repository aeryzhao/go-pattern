package main

import "fmt"

type iSportFactory interface {
	makeShot() iShot
	makeShirt() iShirt
}

func getSportFactory(brand string) (iSportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("nothing same brand")
}
