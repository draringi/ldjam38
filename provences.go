package main

import "math/rand"

const (
	BaseMetalStart = 10000
	BaseMetalVariance = 20000
	BasePopStart = 150000
	BasePopVariance = 200000
)

func RandomMetalStart(wealthFactor float64) int {
	baseMetal := int(wealthFactor * BaseMetalStart)
	variance := int(wealthFactor * BaseMetalVariance)
	if variance == 0 {
		// Set Min Variance
		variance = 100
	}
	return baseMetal + rand.Intn(variance)
}

func RandomPopStart(sizeFactor float64) int {
	basePop := int(sizeFactor * BasePopStart)
	variance := int(sizeFactor * BasePopVariance)
	if variance == 0 {
		// Set Min Variance
		variance = 100
	}
	return basePop + rand.Intn(variance)
}

func ProvenceGLG() {
	p := NewProvence("GLG", 255, 20, 20, "Gilgamesh", RandomMetalStart(25), RandomPopStart(30), LevelCapital)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1575, 1625, 1675, 1675, 1625, 1575, 1525, 1525}
	poly.vy = []int16{1125, 1125, 1175, 1225, 1275, 1275, 1225, 1175}
	poly.namex = 1575;
	poly.namey = 1175;
	polygonMap[p.Tag] = poly
}

func ProvenceQLL() {
	p := NewProvence("QLL", 0, 0, 255, "Quill", RandomMetalStart(10), RandomPopStart(5), LevelTown)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1405, 1218, 1218, 1405}
	poly.vy = []int16{220, 277, 377, 320}
	poly.namex = 1228;
	poly.namey = 300;
	polygonMap[p.Tag] = poly
}

func ProvenceRNK() {
	p := NewProvence("RNK", 5, 240, 30, "Ragnarok", RandomMetalStart(9), RandomPopStart(15), LevelCapital)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1405, 1795, 1795, 1405}
	poly.vy = []int16{220, 220, 450, 450}
	poly.namex = 1415;
	poly.namey = 280;
	polygonMap[p.Tag] = poly
}

func ProvenceWTM() {
	p := NewProvence("WTM", 5, 100, 255, "Watem", RandomMetalStart(20), RandomPopStart(0.5), LevelVillage)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1795, 1795, 2075}
	poly.vy = []int16{220, 320, 320}
	poly.namex = 1800;
	poly.namey = 280;
	polygonMap[p.Tag] = poly
}

func ProvenceENK() {
	p := NewProvence("ENK", 150, 50, 255, "Enkidu", RandomMetalStart(5), RandomPopStart(15), LevelCity)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1575, 1625, 1625, 1575}
	poly.vy = []int16{1125, 1125, 1000, 1000}
	poly.namex = 1590;
	poly.namey = 1020;
	polygonMap[p.Tag] = poly
}

func ProvenceURK() {
	p := NewProvence("URK", 150, 50, 255, "Uruk", RandomMetalStart(2), RandomPopStart(15), LevelCity)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1525, 1525, 1400, 1400}
	poly.vy = []int16{1225, 1175, 1175, 1225}
	poly.namex = 1420;
	poly.namey = 1200;
	polygonMap[p.Tag] = poly
}

func ProvenceNNL() {
	p := NewProvence("NNL", 150, 50, 255, "Ninlil", RandomMetalStart(2), RandomPopStart(15), LevelCity)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1675, 1675, 1800, 1800}
	poly.vy = []int16{1225, 1175, 1175, 1225}
	poly.namex = 1700;
	poly.namey = 1200;
	polygonMap[p.Tag] = poly
}

func ProvenceAN() {
	p := NewProvence("AN", 150, 50, 255, "An", RandomMetalStart(0.1), RandomPopStart(15), LevelCity)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1625, 1575, 1575, 1625}
	poly.vy = []int16{1275, 1275, 1400, 1400}
	poly.namex = 1580;
	poly.namey = 1300;
	polygonMap[p.Tag] = poly
}

func ProvenceTIG() {
	p := NewProvence("TIG", 150, 50, 80, "Tigris", RandomMetalStart(3), RandomPopStart(10), LevelTown)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1400, 1525, 1575, 1575, 1400}
	poly.vy = []int16{1225, 1225, 1275, 1400, 1400}
	poly.namex = 1420;
	poly.namey = 1300;
	polygonMap[p.Tag] = poly
}

func ProvenceEUP() {
	p := NewProvence("EUP", 150, 50, 80, "Euphrates", RandomMetalStart(3), RandomPopStart(10), LevelTown)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1400, 1525, 1575, 1575, 1400}
	poly.vy = []int16{1175, 1175, 1125, 1000, 1000}
	poly.namex = 1420;
	poly.namey = 1100;
	polygonMap[p.Tag] = poly
}

func ProvenceEAM() {
	p := NewProvence("EAM", 150, 50, 80, "Elam", RandomMetalStart(1), RandomPopStart(9), LevelTown)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1800, 1675, 1625, 1625, 1800}
	poly.vy = []int16{1175, 1175, 1125, 1000, 1000}
	poly.namex = 1700;
	poly.namey = 1100;
	polygonMap[p.Tag] = poly
}

func ProvenceEAN() {
	p := NewProvence("EAN", 150, 50, 80, "Eanna", RandomMetalStart(0.6), RandomPopStart(13), LevelTown)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1800, 1675, 1625, 1625, 1800}
	poly.vy = []int16{1225, 1225, 1275, 1400, 1400}
	poly.namex = 1700;
	poly.namey = 1300;
	polygonMap[p.Tag] = poly
}

func ProvenceARB() {
	p := NewProvence("ARB", 0, 128, 128, "Arumba", RandomMetalStart(5), RandomPopStart(8), LevelCity)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1625, 1575, 1575, 1405, 1795, 1625, 1625}
	poly.vy = []int16{1000, 1000, 750, 450, 450, 750, 1000}
	poly.namex = 1500;
	poly.namey = 500;
	polygonMap[p.Tag] = poly
}

func ProvenceBRL() {
	p := NewProvence("BRL", 0, 255, 128, "Bussels", RandomMetalStart(5), RandomPopStart(10), LevelCapital)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1405, 1795, 1795, 1600, 1405}
	poly.vy = []int16{2000, 2000, 2180, 2200, 2180}
	poly.namex = 1500;
	poly.namey = 2050;
	polygonMap[p.Tag] = poly
}

func ProvenceFLA() {
	p := NewProvence("FLA",  255, 100, 5, "Flanders", RandomMetalStart(5), RandomPopStart(8), LevelCity)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1405, 1795, 1625, 1625, 1575, 1575}
	poly.vy = []int16{2000, 2000, 1700, 1400, 1400, 1700}
	poly.namex = 1500;
	poly.namey = 1900;
	polygonMap[p.Tag] = poly
}

func ProvenceBBB() {
	p := NewProvence("BBB", 0, 0, 255, "Bluey", RandomMetalStart(20), RandomPopStart(0.5), LevelVillage)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1405, 1164, 1405}
	poly.vy = []int16{2000, 2100, 2180}
	poly.namex = 1200;
	poly.namey = 2090;
	polygonMap[p.Tag] = poly
}

func ProvenceSCO() {
	p := NewProvence("SCO", 255, 32, 128, "Scotty", RandomMetalStart(10), RandomPopStart(5), LevelTown)
	Provences[p.Tag] = p
	poly := new(polygon)
	poly.vx = []int16{1795, 1795, 1912, 1912}
	poly.vy = []int16{2000, 2180, 2150, 2000}
	poly.namex = 1800;
	poly.namey = 2060;
	polygonMap[p.Tag] = poly
}

func InitProvences() {
	Provences = make(map[string]*Provence)
	polygonMap = make(map[string]*polygon)
	ProvenceGLG()
	ProvenceQLL()
	ProvenceRNK()
	ProvenceWTM()
	ProvenceENK()
	ProvenceURK()
	ProvenceNNL()
	ProvenceAN()
	ProvenceTIG()
	ProvenceEUP()
	ProvenceEAM()
	ProvenceEAN()
	ProvenceARB()
	ProvenceBRL()
	ProvenceFLA()
	ProvenceBBB()
	ProvenceSCO()
}