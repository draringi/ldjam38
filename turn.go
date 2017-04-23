package main

type TurnState uint8

const (
	TurnWaiting TurnState = iota
	TurnFarm
	TurnMine
	TurnTrade
	TurnOrders
	TurnEat
)

type TurnManager struct {
	players uint8
	turnCounter int
	state TurnState
}