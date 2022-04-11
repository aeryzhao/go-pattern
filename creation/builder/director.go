package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setDirector(b iBuilder) {
	d.builder = b
}

func (d *director) builderHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setFloor()
	return d.builder.getHouse()
}
