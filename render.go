package main

import (
	"strconv"
	"log"
	"github.com/veandco/go-sdl2/sdl"
	gfx "github.com/veandco/go-sdl2/sdl_gfx"
	"fmt"
)

const (
	WindowWidth = 800
	WindowHeight = 600
	WorldWidth = 3200
	WorldHeight = 2400
	MaxZoom = 8
	MinZoom = 1
	DefaultZoom = 2
	DefaultX = 0
	DefaultY = 0
	ZoomStep = 0.25
	DiskWorldRadius = 1000
	DiskWorldX = WorldWidth/2
	DiskWorldY = WorldHeight/2
	FontSize = 8
	CameraMoveSpeed = 100
	ProvenceMenuWidth = 320
	ProvenceMenuHeight = 200
	TopBarHeight = 26
)

type Camera struct {
	x int16
	y int16
	ZoomLevel int8
}

func (c *Camera) MaxX() int16 {
	zoom := float64(c.ZoomLevel) * ZoomStep
	cameraWidth := int16(WindowWidth/zoom)
	return WorldWidth-cameraWidth
}

func (c *Camera) MaxY() int16 {
	zoom := float64(c.ZoomLevel) * ZoomStep
	cameraHeight := int16(WindowHeight/zoom)
	return WorldHeight-cameraHeight
}

func (c *Camera) ZoomIn() bool {
	if c.ZoomLevel < MaxZoom {
		c.ZoomLevel++
		return true
	}
	return false
}

func (c *Camera) ZoomOut() bool {
	if c.ZoomLevel > MinZoom {
		c.ZoomLevel--
		return true
	}
	return false
}

func (c *Camera) correct() {
	maxX := c.MaxX()
	if c.x > maxX {
		c.x = maxX
	}
	maxY := c.MaxY()
	if c.y > maxY {
		c.y = maxY
	}
	if c.x < 0 {
		c.x = 0
	}
	if c.y < 0 {
		c.y = 0
	}
}

func (c *Camera) ConvertXY(x, y int16) (int16, int16) {
	zoom := float64(c.ZoomLevel) * ZoomStep
	
	xScaled := int16(zoom*float64(x - c.x))
	yScaled := int16(zoom*float64(y - c.y))
	return xScaled, yScaled
}

func (c *Camera) ConvertR(r int16) int16 {
	zoom := float64(c.ZoomLevel) * ZoomStep
	return int16 (zoom * float64(r))
}

func (c *Camera) ConvertToWorld(x, y int32) (int32, int32) {
	zoom := float64(c.ZoomLevel) * ZoomStep
	
	xScaled := int32(float64(x)/zoom) + int32(c.x)
	yScaled := int32(float64(y)/zoom) + int32(c.y)
	return xScaled, yScaled
}

func NewCamera() *Camera {
	c := new(Camera)
	c.x = DefaultX
	c.y = DefaultY
	c.ZoomLevel = DefaultZoom
	return c
}

type polygon struct {
	vx []int16
	vy []int16
	namex int16
	namey int16
}

var polygonMap map[string]*polygon

func drawProvence(p *Provence, r *sdl.Renderer) error {
	poly := polygonMap[p.Tag]
	if poly == nil {
		return fmt.Errorf("Provence tag \"%s\" not in polygon map", p.Tag)
	}
	vertexCount := len(poly.vx)
	vx := make([]int16, vertexCount)
	vy := make([]int16, vertexCount)
	for i, _ := range(vx) {
		vx[i], vy[i] = GameCam.ConvertXY(poly.vx[i], poly.vy[i])
	}
	ok := gfx.FilledPolygonColor(r, vx, vy, p.Colour)
	if !ok {
		return fmt.Errorf("Error drawing provence \"%s\"", p.Tag)
	}
	if GameCam.ZoomLevel > 2 {
		nameX, nameY := GameCam.ConvertXY(poly.namex, poly.namey)
		ok = gfx.StringColor(r, int(nameX), int(nameY), p.Name, colourProvenceName)
		if !ok {
			return fmt.Errorf("Error drawing name for provence \"%s\"", p.Tag)
		}
		cx, cy := p.Centre()
		symX, symY := GameCam.ConvertXY(cx, cy)
		var colour sdl.Color
		if p.Owner != nil {
			colour = p.Owner.Colour
		} else {
			colour = colourBlack
		}
		sym := p.Level.Symbol()
		if sym == nil {
			log.Panicf("No symbol found for provence level \"%v\" (%d)", p.Level, int(p.Level))
		}
		symX -= int16(sym.Width/2)
		symY -= int16(sym.Height/2)
		box := sdl.Rect{int32(symX), int32(symY), int32(sym.Width), int32(sym.Height)}
		r.SetDrawColor(colour.R, colour.G, colour.B, colour.A)
		r.FillRect(&box)
		r.Copy(sym.Texture, nil, &box)
	}
	return nil
}

func drawWorld(r *sdl.Renderer) error {
	r.SetDrawColor(0,0,0,255)
	r.Clear()
	circleX, circleY := GameCam.ConvertXY(DiskWorldX, DiskWorldY)
	circleR := GameCam.ConvertR(DiskWorldRadius)
	gfx.FilledCircleColor(r, int(circleX), int(circleY), int(circleR), colourWorldBackground)
	for _, p := range(Provences) {
		err := drawProvence(p, r)
		if err != nil {
			log.Panicln(err)
		}
	}
	//gfx.CircleColor(r, int(circleX), int(circleY), int(circleR), colourRed)
	drawButton(r, ButtonEndTurn, WindowWidth-ButtonCentreX, WindowHeight-ButtonCentreY)
	return nil
}

func drawButton(r *sdl.Renderer, ID ButtonID, x, y int32) {
	buttonRect := sdl.Rect{0, 0, ButtonWidth, ButtonHeight}
	x_mod := x - ButtonWidth/2
	y_mod := y - ButtonHeight/2
	targetRect := sdl.Rect{x_mod, y_mod, ButtonWidth, ButtonHeight}
	button := buttons[ID]
	if button == nil {
		log.Panicf("Invalid Button ID %d\n", int(ID))
	}
	r.Copy(button, &buttonRect, &targetRect)
}

func drawMenu(r *sdl.Renderer) {
	MenuTop := int32(WindowHeight/2-MenuHeight/2)
	menuBackground := sdl.Rect{WindowWidth/2-MenuWidth/2, MenuTop, MenuWidth, MenuHeight}
	r.SetDrawColor(colourMenuBackground.R, colourMenuBackground.G, colourMenuBackground.B, colourMenuBackground.A)
	r.FillRect(&menuBackground)
	y := MenuTop + 2*ButtonHeight
	drawButton(r, ButtonResume, WindowWidth/2, y)
	y += 2*ButtonHeight
	drawButton(r, ButtonNewGame, WindowWidth/2, y)
	y += 2*ButtonHeight
	drawButton(r, ButtonQuit, WindowWidth/2, y)
}

const (
	TopBarValueOffset = 10
	TopBarSeperator = 50
	TopBarTextOffset = 3
)

func drawTopMenu(r *sdl.Renderer) {
	PlayerNation := Nations["PLA"]
	if PlayerNation == nil {
		return
	}
	topBarBackground := sdl.Rect{0, 0, WindowWidth, TopBarHeight}
	r.SetDrawColor(colourTopBarBackground.R, colourTopBarBackground.G, colourTopBarBackground.B, colourTopBarBackground.A)
	r.FillRect(&topBarBackground)
	x := int32(15)
	var sym *Symbol
	var drawRect sdl.Rect
	sym = SymbolMap[SymbolFood]
	drawRect = sdl.Rect{x, int32(TopBarHeight/2-sym.Height/2), int32(sym.Width), int32(sym.Height) }
	r.Copy(sym.Texture, nil, &drawRect)
	x  += sym.Width + TopBarValueOffset
	gfx.StringColor(r, int(x), TopBarHeight/2 - TopBarTextOffset, strconv.Itoa(PlayerNation.Food), colourWhite)
	x += TopBarSeperator
	sym = SymbolMap[SymbolMetal]
	drawRect = sdl.Rect{x, int32(TopBarHeight/2-sym.Height/2), int32(sym.Width), int32(sym.Height) }
	r.Copy(sym.Texture, nil, &drawRect)
	x  += sym.Width + TopBarValueOffset
	gfx.StringColor(r, int(x), TopBarHeight/2 - TopBarTextOffset, strconv.Itoa(PlayerNation.Metal), colourWhite)
	x += TopBarSeperator
	x += TopBarSeperator
	sym = SymbolMap[SymbolSword]
	drawRect = sdl.Rect{x, int32(TopBarHeight/2-sym.Height/2), int32(sym.Width), int32(sym.Height) }
	r.Copy(sym.Texture, nil, &drawRect)
	x  += sym.Width + TopBarValueOffset
	gfx.StringColor(r, int(x), TopBarHeight/2 - TopBarTextOffset, strconv.Itoa(PlayerNation.Equipment[Sword]), colourWhite)
	x += TopBarSeperator
	sym = SymbolMap[SymbolPike]
	drawRect = sdl.Rect{x, int32(TopBarHeight/2-sym.Height/2), int32(sym.Width), int32(sym.Height) }
	r.Copy(sym.Texture, nil, &drawRect)
	x  += sym.Width + TopBarValueOffset
	gfx.StringColor(r, int(x), TopBarHeight/2 - TopBarTextOffset, strconv.Itoa(PlayerNation.Equipment[Pike]), colourWhite)
	x += TopBarSeperator
	sym = SymbolMap[SymbolCrossbow]
	drawRect = sdl.Rect{x, int32(TopBarHeight/2-sym.Height/2), int32(sym.Width), int32(sym.Height) }
	r.Copy(sym.Texture, nil, &drawRect)
	x  += sym.Width + TopBarValueOffset
	gfx.StringColor(r, int(x), TopBarHeight/2 - TopBarTextOffset, strconv.Itoa(PlayerNation.Equipment[Crossbow]), colourWhite)
	x += TopBarSeperator
	sym = SymbolMap[SymbolBow]
	drawRect = sdl.Rect{x, int32(TopBarHeight/2-sym.Height/2), int32(sym.Width), int32(sym.Height) }
	r.Copy(sym.Texture, nil, &drawRect)
	x  += sym.Width + TopBarValueOffset
	gfx.StringColor(r, int(x), TopBarHeight/2 - TopBarTextOffset, strconv.Itoa(PlayerNation.Equipment[Bow]), colourWhite)
	x += TopBarSeperator
	sym = SymbolMap[SymbolLance]
	drawRect = sdl.Rect{x, int32(TopBarHeight/2-sym.Height/2), int32(sym.Width), int32(sym.Height) }
	r.Copy(sym.Texture, nil, &drawRect)
	x  += sym.Width + TopBarValueOffset
	gfx.StringColor(r, int(x), TopBarHeight/2 - TopBarTextOffset, strconv.Itoa(PlayerNation.Equipment[Lance]), colourWhite)
}

const (
	ProvenceMenuNameOffsetX = 20
	ProvenceMenuNameOffsetY = 15
	ProvenceMenuOwnerOffsetX = 20
	ProvenceMenuOwnerOffsetY = 35
	ProvenceMenuPopOffsetX = 20
	ProvenceMenuPopOffsetY = 45
	ProvenceMenuMetalOffsetX = 20
	ProvenceMenuMetalOffsetY = 55
	ProvenceMenuLevelOffsetX = 20
	ProvenceMenuLevelOffsetY = 25
	ProvenceMenuArmiesOffsetX = 20
	ProvenceMenuArmiesOffsetY = 65
	ProvenceMenuBuildOffsetX = 20
	ProvenceMenuBuildOffsetY = 75
	ProvenceMenuTrainOffsetX = 20
	ProvenceMenuTrainOffsetY = 85
)

func drawProvenceMenu(r *sdl.Renderer) {
	MenuTop := WindowHeight - ProvenceMenuHeight
	r.SetDrawColor(colourProvenceMenuBackground.R, colourProvenceMenuBackground.G, colourProvenceMenuBackground.B, colourProvenceMenuBackground.A)
	menuBackground := sdl.Rect{0, int32(MenuTop), ProvenceMenuWidth, ProvenceMenuHeight}
	r.FillRect(&menuBackground)
	gfx.StringColor(r, ProvenceMenuNameOffsetX, MenuTop + ProvenceMenuNameOffsetY, SelectedProvence.Name, colourWhite)
	popString := fmt.Sprintf("Population: %d", SelectedProvence.Population)
	gfx.StringColor(r, ProvenceMenuPopOffsetX, MenuTop + ProvenceMenuPopOffsetY, popString, colourWhite)
	metalString := fmt.Sprintf("Unmined Metal: %d", SelectedProvence.MetalCount)
	gfx.StringColor(r, ProvenceMenuMetalOffsetX, MenuTop + ProvenceMenuMetalOffsetY, metalString, colourWhite)
	levelString := fmt.Sprintf("Provence Level: %v", SelectedProvence.Level)
	gfx.StringColor(r, ProvenceMenuLevelOffsetX, MenuTop + ProvenceMenuLevelOffsetY, levelString, colourWhite)
	var ownerName string
	if SelectedProvence.Owner != nil {
		ownerName = SelectedProvence.Owner.Name
	} else {
		ownerName = "Unowned"
	}
	ownerString := fmt.Sprintf("Owner: %s", ownerName)
	gfx.StringColor(r, ProvenceMenuOwnerOffsetX, MenuTop + ProvenceMenuOwnerOffsetY, ownerString, colourWhite)
	player := Nations["PLA"]
	if player == nil {
		// stop here
		return
	}
	if player != SelectedProvence.Owner {
		// stop here
		return
	}
	// Terrible efficency, but in a HURRY
	localFriendlies := len(FilterArmiesPresent(player, SelectedProvence))
	armiesString := fmt.Sprintf("Armies: %d", localFriendlies)
	gfx.StringColor(r, ProvenceMenuArmiesOffsetX, MenuTop + ProvenceMenuArmiesOffsetY, armiesString, colourWhite)
	ProvenceOrders := ProvincialOrders(player.OrderQueue, SelectedProvence)
	var trainOrder *TrainOrder
	var buildOrder *BuildOrder
	for _, order := range(ProvenceOrders) {
		switch o := order.(type) {
		case *BuildOrder:
			buildOrder = o
		case *TrainOrder:
			trainOrder = o
		}
	}
	var buildString string
	var trainString string
	if buildOrder != nil {
		buildString = fmt.Sprintf("Building: %v (%d remaining)", buildOrder.itemType, buildOrder.count)
	} else {
		buildString = "Nothing Being Built"
	}
	gfx.StringColor(r, ProvenceMenuBuildOffsetX, MenuTop + ProvenceMenuBuildOffsetY, buildString, colourWhite)
	if trainOrder != nil {
		trainString = fmt.Sprintf("TTraining: %v (%d remaining)", trainOrder.unitType, buildOrder.count)
	} else {
		trainString = "No troops being trained"
	}
	gfx.StringColor(r, ProvenceMenuTrainOffsetX, MenuTop + ProvenceMenuTrainOffsetY, trainString, colourWhite)
	drawButton(r, ButtonArmies, ButtonArmiesX + ButtonWidth/2, ButtonArmiesY + ButtonHeight/2)
	drawButton(r, ButtonBuild, ButtonBuildX + ButtonWidth/2, ButtonBuildY + ButtonHeight/2)
	drawButton(r, ButtonTrain, ButtonTrainX + ButtonWidth/2, ButtonTrainY + ButtonHeight/2)
}