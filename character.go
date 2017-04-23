package main

type CharPortrait uint8
type CharStatus uint8

const (
	Healthy CharStatus = iota
	Sick
	Injured
	Dead
)

type Character struct {
	Name string
	Age int
	Gender string
	Status CharStatus
	BetterIn int
	//Race *Race
	Partner *Character
	Children []*Character
	Employee *Nation
	Portrait CharPortrait
}