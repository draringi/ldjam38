package main

import (
	"math/rand"
	"sync"
)

type TurnState uint8

const (
	TurnWaiting TurnState = iota
	TurnFarm
	TurnMine
	TurnTrade
	TurnOrders
	TurnEat
)

type PlayerTurnStatus struct {
	submitted bool
	orders []Order
	index int
	foodNeeded int
	modifier float64
	nation *Nation
}

type TurnManager struct {
	lock *sync.Mutex
	turnCounter int
	state TurnState
	status map[string]*PlayerTurnStatus
}

func NewTurnManager() *TurnManager {
	tm := new(TurnManager)
	tm.turnCounter = 1
	tm.state = TurnWaiting
	tm.status = make(map[string]*PlayerTurnStatus)
	for _, nation := range(Nations) {
		tm.status[nation.Tag] = new(PlayerTurnStatus)
		tm.status[nation.Tag].nation = nation
	}
	tm.lock = new(sync.Mutex)
	return tm
}

func (tm *TurnManager) SubmitTurn(nation *Nation, orders []Order) {
	tm.lock.Lock()
	tm.status[nation.Tag].submitted = true
	tm.status[nation.Tag].index = 0
	tm.status[nation.Tag].orders = orders
	tm.lock.Unlock()
}

func (tm *TurnManager) PlayersRemaining() int {
	if tm.state != TurnWaiting {
		return 0
	}
	count := 0
	tm.lock.Lock()
	for _, n := range(tm.status) {
		if !n.submitted {
			count++
		}
	}
	tm.lock.Unlock()
	return count
}

func (tm *TurnManager) Execute(gState *GameState) {
	tm.state = TurnEat
	nationsWithOrders := make([]*PlayerTurnStatus, len(tm.status))
	i := 0
	for _, n := range(tm.status) {
		n.index = 0
		n.foodNeeded = 0
		nationsWithOrders[i] = n
		i++
	}
	
	for _, p := range(Provences) {
		n := p.Owner
		if n != nil {
			n.Metal += p.MineMetal()
			n.Metal += p.PerformTrade()
			n.Food += p.GrowFood()
			tm.status[n.Tag].foodNeeded += p.EatFood()
		}
	}
	for len(nationsWithOrders) > 0 {
		i := rand.Intn(len(nationsWithOrders))
		n := nationsWithOrders[i]
		if n.index == len(n.orders) {
			n.index = 0
			if i == 0 {
				nationsWithOrders = nationsWithOrders[1:]
			} else if i + 1 == len(nationsWithOrders) {
				nationsWithOrders = nationsWithOrders[:i]
			} else {
				nationsWithOrders = append(nationsWithOrders[:i], nationsWithOrders[i+1:]...)
			}
			continue
		}
		n.orders[n.index].Execute(n.nation)
		n.index++
	}
	for _, n := range(tm.status) {
		if n.foodNeeded > n.nation.Food {
			foodProvided := float64(n.nation.Food)/float64(n.foodNeeded)
			n.nation.Food = 0
			n.modifier = (0.8 - foodProvided)/0.8
		} else {
			n.nation.Food -= n.foodNeeded
			n.modifier = 1
		}
	}
	for _, p := range(Provences) {
		n := p.Owner
		if n != nil {
			modifier := tm.status[n.Tag].modifier
			p.GrowPop(modifier)
		} else {
			p.GrowPop(0.25)
		}
	}
	for _, n := range(tm.status) {
		var newOrders []Order
		for _, o := range n.orders {
			if !o.IsDone() {
				newOrders = append(newOrders, o)
			}
			n.nation.OrderQueue = newOrders
		}
		n.submitted = false
	}
	tm.state = TurnWaiting
	tm.turnCounter++
	*gState = GameNewTurn
}

var (
	ActiveTurnManager *TurnManager
)