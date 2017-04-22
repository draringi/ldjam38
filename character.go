package main

type CharPortrait uint8

type Character struct {
	Name string
	Age int
	Gender string
	//Race *Race
	Partner *Character
	Children []*Character
	Employee *Nation
	Portrait CharPortrait
}