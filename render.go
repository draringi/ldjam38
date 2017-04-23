package main

import (
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
	return nil
}