package main

type Army struct {
	General *Character
	Owner *Nation
	Location *Provence
	Morale float64
	Troops map[UnitType]int
	
}

func NewArmy(n *Nation, l *Provence) *Army {
	a := new(Army)
	a.Owner = n
	a.Location = l
	a.Troops = make(map[UnitType]int)
	return a
}