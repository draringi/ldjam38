package main

type Army struct {
	General *Character
	Owner *Nation
	Location *Provence
}

func NewArmy(n *Nation, l *Provence) *Army {
	a := new(Army)
	a.Owner = n
	a.Location = l
}