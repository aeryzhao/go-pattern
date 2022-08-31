package main

import "fmt"

type house struct {
	doorType   string
	windowType string
	floor      int
}

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

type iglooBuilder struct {
	doorType   string
	windowType string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Igloo Window Type"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Igloo Door Type"
}

func (b *iglooBuilder) setFloor() {
	b.floor = 2
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type normalBuilder struct {
	doorType   string
	windowType string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Normal door Type"
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Normal Window Type"
}

func (b *normalBuilder) setFloor() {
	b.floor = 1
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) builderHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")

	director := newDirector(normalBuilder)
	normalHouse := director.builderHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)
}
