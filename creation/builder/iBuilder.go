package main

type iBuilder interface {
	setDoorType()
	setWindowType()
	setFloor()
	getHouse() house
}

func getBuilder(s string) iBuilder {
	if s == "normal" {
		return newNormalBuilder()
	}
	if s == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
