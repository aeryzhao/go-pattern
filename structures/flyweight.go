package main

import "fmt"

const (
	TerroristDressType         = "tDress"
	CounterTerroristDressType  = "ctDress"
	TerroristPlayerType        = "T"
	CounterTerroristPlayerType = "CT"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) dress {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType]
	}
	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = &terroristDress{
			color: "red",
		}
	case CounterTerroristDressType:
		d.dressMap[dressType] = &counterTerroristDress{
			color: "green",
		}
	}
	return d.dressMap[dressType]
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

type dress interface {
	getColor() string
}

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{
		color: "red",
	}
}

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{
		color: "green",
	}
}

type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType string, dressType string) *player {
	return &player{
		playerType: playerType,
		dress:      getDressFactorySingleInstance().getDressByType(dressType),
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	c.terrorists = append(c.terrorists, newPlayer(TerroristPlayerType, dressType))
}

func (c *game) addCounterTerrorist(dressType string) {
	c.counterTerrorists = append(c.counterTerrorists, newPlayer(CounterTerroristPlayerType, dressType))
}

func main() {
	g := newGame()

	g.addTerrorist(TerroristDressType)
	g.addTerrorist(TerroristDressType)
	g.addTerrorist(TerroristDressType)
	g.addTerrorist(TerroristDressType)

	g.addCounterTerrorist(CounterTerroristDressType)
	g.addCounterTerrorist(CounterTerroristDressType)
	g.addCounterTerrorist(CounterTerroristDressType)
	g.addCounterTerrorist(CounterTerroristDressType)

	for dressType, dress := range getDressFactorySingleInstance().dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
