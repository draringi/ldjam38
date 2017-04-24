package main

type Army struct {
	General *Character
	Owner *Nation
	Location *Provence
	Name string
	Morale float64
	Troops map[UnitType]int
	
}

func NewArmy(n *Nation, l *Provence, name string) *Army {
	a := new(Army)
	a.Owner = n
	a.Location = l
	a.Troops = make(map[UnitType]int)
	a.Name = name
	n.Armies = append(n.Armies, a)
	return a
}

func FilterArmiesPresent(n *Nation, p *Provence) []*Army {
	var armies []*Army
	for _, army := range(n.Armies) {
		if army.Location == p {
			armies = append(armies, army)
		}
	}
	return armies
}

var (
	SelectedArmy *Army
)