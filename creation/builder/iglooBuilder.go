package main

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
