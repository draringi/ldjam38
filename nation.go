package main

import "github.com/veandco/go-sdl2/sdl"

var Nations map[string]*Nation

type Nation struct {
	Tag string
	Name string
	Colour sdl.Color
	Metal int
	Food int
	Leader *Character
	Equipment map[ItemType]int
	OrderQueue []Order
	Armies []*Army
}

const (
	NationStartingFood = 10000
	NationStartingMetal = 10000
)

func NewNation(name, tag string, leader *Character, red, green, blue uint8) *Nation {
	n := new(Nation)
	n.Name = name
	n.Tag = tag
	n.Colour = sdl.Color{red, green, blue, 255}
	n.Leader = leader
	n.Metal = NationStartingMetal
	n.Food = NationStartingFood
	n.Equipment = make(map[ItemType]int)
	return n
}

func (n *Nation) EndTurn() {
	ActiveTurnManager.SubmitTurn(n, n.OrderQueue)
}
