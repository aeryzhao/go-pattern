package main

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
