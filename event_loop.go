package main

import (
	"github.com/veandco/go-sdl2/sdl"
 	gfx "github.com/veandco/go-sdl2/sdl_gfx"
)

func game_loop(r *sdl.Renderer) {
	running := true
	moveRight := false
	moveLeft := false
	moveUp := false
	moveDown := false
	var fpsManager gfx.FPSmanager
	gfx.InitFramerate(&fpsManager)
	for running{
		zoomed := false
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseWheelEvent:
				if !zoomed {
					if t.Y > 0 {
						zoomed = GameCam.ZoomIn()
					} else if t.Y < 0 && GameCam.ZoomLevel > MinZoom {
						zoomed = GameCam.ZoomOut()
					}
				}
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_KP_PLUS, sdl.K_PLUS, sdl.K_PAGEUP:
					if !zoomed {
						zoomed = GameCam.ZoomIn()
					}
				case sdl.K_KP_MINUS, sdl.K_MINUS, sdl.K_PAGEDOWN:
					if !zoomed {
						zoomed = GameCam.ZoomOut()
					}
				case sdl.K_RIGHT:
					moveRight = true
				case sdl.K_LEFT:
					moveLeft = true
				case sdl.K_UP:
					moveUp = true
				case sdl.K_DOWN:
					moveDown = true
				case sdl.K_ESCAPE:
					running = false
				}
			case *sdl.KeyUpEvent:
				switch t.Keysym.Sym {
				case sdl.K_RIGHT:
					moveRight = false
				case sdl.K_LEFT:
					moveLeft = false
				case sdl.K_UP:
					moveUp = false
				case sdl.K_DOWN:
					moveDown = false
				}
			}
		}
		var moveAmount int16
		if moveUp || moveDown || moveLeft || moveRight {
			moveAmount = int16(CameraMoveSpeed/(float64(GameCam.ZoomLevel)*ZoomStep))
		}
		if moveUp {
			GameCam.y -= moveAmount
		}
		if moveDown {
			GameCam.y += moveAmount
		}
		if moveLeft {
			GameCam.x -= moveAmount
		}
		if moveRight {
			GameCam.x += moveAmount
		}
		GameCam.correct()
		err := drawWorld(r)
		if err != nil {
			panic(err)
		}
		
		r.Present()
		gfx.FramerateDelay(&fpsManager)
	}
}