package main

import "github.com/veandco/go-sdl2/sdl"

type nodeLevel int

const (
	LevelVillage nodeLevel = iota
	LevelTown
	LevelCity
	LevelCapital
)

const (
	PopWorkerModifer = 10000
	PopEatModifer = 7777
	PopBaseGrowthRate = 0.04
	MetalMineRate = 3000
)

type Provence struct {
	Tag string
	Colour sdl.Color
	Name string
	DefendabilityMod int
	FoodGrowMod int
	MineMod int
	TradeMod int
	Level nodeLevel
	Owner *Nation
	Population int
	MetalCount int
	Neighbours []*Provence
}

func NewProvence(tag string, red, green, blue uint8, name string, metal, startPop int, startLevel nodeLevel) *Provence {
	p := new(Provence)
	p.Tag = tag
	p.Name = name
	p.MetalCount = metal
	p.Population = startPop
	p.Level = startLevel
	p.Colour = sdl.Color{red, green, blue, 0}
	return p
}

// MineMetal calculates how much metal to mine in a provence, removes it from the provence metal count,
// and returns it.
func (p *Provence) MineMetal() int {
	level := int(p.Level)
	countModifier := float64(p.MetalCount) / MetalMineRate
	adjustedModifier := float64((4-level)*10 + p.MineMod)
	popModifier := float64(p.Population) / PopWorkerModifer
	MinedMetalAmount := int(countModifier * adjustedModifier * popModifier)
	if MinedMetalAmount > p.MetalCount {
		MinedMetalAmount = p.MetalCount
	}
	p.MetalCount -= MinedMetalAmount
	return MinedMetalAmount
}

func (p *Provence) GrowFood() int {
	level := int(p.Level)
	adjustedModifier := float64((4-level)*10 + p.FoodGrowMod)
	popModifier := float64(p.Population) / PopWorkerModifer
	return int(popModifier * adjustedModifier)
}

func (p *Provence) EatFood() int {
	fPop := float64(p.Population)
	return int(fPop/PopEatModifer)
}

func (p *Provence) PerformTrade() int {
	level := int(p.Level)
	adjustedModifier := float64 (level * 10 + p.TradeMod)
	popModifier := float64(p.Population) / PopWorkerModifer
	return int(popModifier * adjustedModifier)
}

func (p *Provence) GrowPop(foodAvailableModifer float64) {
	fPop := float64(p.Population)
	p.Population += int(fPop * PopBaseGrowthRate * foodAvailableModifer)
	if p.Population <= 0 {
		// Terrible things, but for starters no production
		// No Trade, no defencability. Can it be called owned?
	}
}

func (p *Provence ) Defendability() int {
	return 10*int(p.Level) + p.DefendabilityMod
}

type Path [2]string

var (
	Provences map[string]*Provence
	Paths []Path
)