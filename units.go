package main

type ItemType uint8

const (
	Sword ItemType = iota
	Pike
	Crossbow
	Bow
	Lance
)

var ItemCost = map[ItemType]int {
	Sword: 5,
	Pike: 4,
	Crossbow: 10,
	Bow: 1,
	Lance: 15,
}

var ItemRate = map[ItemType]float64 {
	Sword: 2,
	Pike: 1,
	Crossbow: 2,
	Bow: 2,
	Lance: 5,
}

type UnitType uint8

const (
	Swordsman UnitType = iota
	Pikeman
	Crossbowman
	Archer
	Lancer
)

var UnitCost = map[UnitType]int {
	Swordsman: 10,
	Pikeman: 10,
	Crossbowman: 10,
	Archer: 10,
	Lancer: 10,
}

var UnitRate = map[UnitType]int {
	Swordsman: 1000,
	Pikeman: 1000,
	Crossbowman: 900,
	Archer: 250,
	Lancer: 50,
}