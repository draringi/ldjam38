package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type nodeLevel int

const (
	LevelVillage nodeLevel = iota
	LevelTown
	LevelCity
	LevelCapital
)

func (n nodeLevel) String() string {
	switch n {
	case LevelVillage:
		return "Village"
	case LevelTown:
		return "Town"
	case LevelCity:
		return "City"
	case LevelCapital:
		return "Capital"
	default:
		return "[UNKNOWN]"
	}
}

func (n nodeLevel) Symbol() *Symbol {
	switch n {
	case LevelVillage:
		return SymbolMap[SymbolVillage]
	case LevelTown:
		return SymbolMap[SymbolTown]
	case LevelCity:
		return SymbolMap[SymbolCity]
	case LevelCapital:
		return SymbolMap[SymbolCapital]
	default:
		return nil
	}
}

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
	p.Colour = sdl.Color{red, green, blue, 100}
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

func isOdd(n int) bool {
	n_div := n >> 1
	return (n - (n_div<<1)) == 1
}

func inBounds(n float64, p1, p2 int32) bool {
	if p1 > p2 {
		p1, p2 = p2, p1
	}
	if n >= float64(p1) && n < float64(p2) {
		return true
	}
	return false
}

func (p *Provence) Contains(x, y int32) bool {
	
	poly := polygonMap[p.Tag]
	if poly == nil {
		log.Panicf("Unknown tag %s\n", p.Tag)
	}
	point_count := len(poly.vx)
	poly_points := make([]sdl.Point, point_count)
	for i, _ := range(poly.vx) {
		poly_points[i].X = int32(poly.vx[i])
		poly_points[i].Y = int32(poly.vy[i])
	}
	containingRect, _ := sdl.EnclosePoints(poly_points, nil)
	point := new(sdl.Point)
	point.X = x
	point.Y = y
	if !point.InRect(&containingRect) {
		return false
	}
	poly_lines := make([][2]sdl.Point, point_count)
	for i, p := range(poly_points) {
		poly_lines[i][0] = p
		if i==0 {
			poly_lines[point_count -1][1] = p
		} else {
			poly_lines[i-1][1] = p
		}
	}
	cross_count := 0
	parrallel_count := 0
	for _, line := range(poly_lines) {
		line_y_diff := line[0].Y - line[1].Y
		line_x_diff := line[0].X - line[1].X
		determinate := point.Y*line_x_diff - point.X*line_y_diff
		if determinate == 0 {
			parrallel_count++
			continue
		}
		line_prod := line[0].X * line[1].Y - line[0].Y * line[1].X
		if line_x_diff != 0 {
			cross_x := float64(line_prod * point.X)/float64(determinate)
			if !(inBounds(cross_x, line[0].X, line[1].X) && cross_x < float64(point.X)) {
				continue
			}
		}
		if line_y_diff != 0 {
			cross_y := float64(line_prod * point.Y)/float64(determinate)
			if inBounds(cross_y, line[0].Y, line[1].Y) && cross_y < float64(point.Y) {
				cross_count++
			}
		} else {
			cross_count++
		}
	}
	return isOdd(cross_count)
}

func (p *Provence) Centre() (int16, int16) {
	poly := polygonMap[p.Tag]
	if poly == nil {
		log.Panicf("Unknown tag %s\n", p.Tag)
	}
	centreX := 0
	for _, val := range(poly.vx) {
		centreX += int(val)
	}
	centreX /= len(poly.vx)
	centreY := 0
	for _, val := range(poly.vy) {
		centreY += int(val)
	}
	centreY /= len(poly.vy)
	return int16(centreX), int16(centreY)
}

func GetProvenceAt(x, y int32) *Provence {
	for _, prov := range(Provences) {
		if prov.Contains(x, y) {
			return prov
		}
	}
	return nil
}

type Path [2]string

var (
	Provences map[string]*Provence
	SelectedProvence *Provence = nil
	Paths []Path
)