package main

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
 	gfx "github.com/veandco/go-sdl2/sdl_gfx"
)

type GameState int

const (
	GamePlay GameState = iota
	GameMenu
	GameHelp
	GameQuit
	GameReset
	GameEoT
	GameBuild
	GameTrain
	GameNewTurn
)

func game_loop(r *sdl.Renderer) {
	moveRight := false
	moveLeft := false
	moveUp := false
	moveDown := false
	state := GameReset
	var fpsManager gfx.FPSmanager
	gfx.InitFramerate(&fpsManager)
	for state != GameQuit {
		if state == GameReset {
			GameCam = NewCamera()
			SelectedArmy = nil
			SelectedProvence = nil
			InitProvences()
			InitNations()
			state = GameNewTurn
			ActiveTurnManager = NewTurnManager()
		}
		if state == GameNewTurn {
			for _, nation := range(Nations) {
				if nation.Tag != "PLA" {
					go RunAI(nation, ActiveTurnManager)
				}
			}
			state = GamePlay
			log.Printf("Start of Turn: %d\n", ActiveTurnManager.turnCounter)
		}
		zoomed := false
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				state = GameQuit
			case *sdl.MouseWheelEvent:
				if !zoomed && state == GamePlay {
					if t.Y > 0 {
						zoomed = GameCam.ZoomIn()
					} else if t.Y < 0 && GameCam.ZoomLevel > MinZoom {
						zoomed = GameCam.ZoomOut()
					}
				}
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_KP_PLUS, sdl.K_PLUS, sdl.K_PAGEUP:
					if !zoomed && state == GamePlay{
						zoomed = GameCam.ZoomIn()
					}
				case sdl.K_KP_MINUS, sdl.K_MINUS, sdl.K_PAGEDOWN:
					if !zoomed && state == GamePlay{
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
					switch state {
					case GamePlay:
						state = GameMenu
					case GameMenu:
						state = GamePlay
					}
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
			case *sdl.MouseButtonEvent:
				if t.Type == sdl.MOUSEBUTTONDOWN && t.Button == sdl.BUTTON_LEFT {
					if CheckButtonPress(t.X, t.Y, &state) || state != GamePlay {
						continue
					}
					if SelectedProvence != nil {
						if CheckProvenceMenu(t.X, t.Y, &state) {
							continue
						}
					}
					worldX, worldY := GameCam.ConvertToWorld(t.X, t.Y)
					prov := GetProvenceAt(worldX, worldY)
					if prov != nil {
						log.Printf("Clicked on Provence %s\n", prov.Tag)
					}
					SelectedProvence = prov
				}
			}
		}
		if state == GamePlay {
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
		}
		GameCam.correct()
		err := drawWorld(r)
		if err != nil {
			panic(err)
		}
		if SelectedProvence != nil {
			drawProvenceMenu(r)
		}
		drawTopMenu(r)
		if state == GameMenu {
			drawMenu(r)
		}
		
		if state == GameEoT {
			if ActiveTurnManager.PlayersRemaining() == 0 {
				ActiveTurnManager.Execute(&state)
			}
		}
		
		r.Present()
		gfx.FramerateDelay(&fpsManager)
	}
}