package main

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	//foodString = "🍞"
	foodString = "F"
	metalString = "M"
	//bowString = "🏹"
	bowString = "b"
	swordString = "s"
	pikeString = "p"
	crossbowString = "c"
	lanceString = "l"
	//villageString = "🐷"
	villageString = "v"
	//townString = "🏠"
	townString = "t"
	//cityString = "🏰"
	cityString = "c"
	//capitalString = "👑"
	capitalString = "C"
)

type SymbolID int

const (
	SymbolFood SymbolID = iota
	SymbolMetal
	SymbolBow
	SymbolSword
	SymbolPike
	SymbolCrossbow
	SymbolLance
	SymbolCapital	
	SymbolCity
	SymbolTown
	SymbolVillage
	
)

type Symbol struct {
	ID SymbolID
	Texture *sdl.Texture
	Width int32
	Height int32
}

func MakeSymbol(r *sdl.Renderer, id SymbolID, str string) {
	sym := new(Symbol)
	sym.ID = id
	text, err := DroidSansMono.RenderUTF8_Blended(str, colourWhite)
	if err != nil {
		log.Panicln(err)
	}
	sym.Width = text.W
	sym.Height = text.H
	sym.Texture, err = r.CreateTextureFromSurface(text)
	if err != nil {
		log.Panicln(err)
	}
	text.Free()
	SymbolMap[id] = sym
}

var SymbolMap map[SymbolID]*Symbol

// Must come after LoadButtons
func LoadSymbols(r *sdl.Renderer) {
	SymbolMap = make(map[SymbolID]*Symbol)
	MakeSymbol(r, SymbolCapital, capitalString)
	MakeSymbol(r, SymbolCity, cityString)
	MakeSymbol(r, SymbolTown, townString)
	MakeSymbol(r, SymbolVillage, villageString)
	MakeSymbol(r, SymbolFood, foodString)
	MakeSymbol(r, SymbolMetal, metalString)
	MakeSymbol(r, SymbolBow, bowString)
	MakeSymbol(r, SymbolSword, swordString)
	MakeSymbol(r, SymbolPike, pikeString)
	MakeSymbol(r, SymbolCrossbow, crossbowString)
	MakeSymbol(r, SymbolLance, lanceString)
}

func CleanupSymbols() {
	for _, sym := range(SymbolMap) {
		sym.Texture.Destroy()
	}
	SymbolMap = nil
}