package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const (
	GameName = "It's a Small World Afterall"
)

var GameCam *Camera

func main() {
		sdl.Init(sdl.INIT_VIDEO|sdl.INIT_TIMER|sdl.INIT_AUDIO)
		defer sdl.Quit()
		var sdlVersion sdl.Version
		sdl.GetVersion(&sdlVersion)
		log.Printf("SDL Initialized. Version: %d.%d.%d\n", sdlVersion.Major, sdlVersion.Minor, sdlVersion.Patch)
		log.Printf("CPU Count: %d, OS: %s\n", sdl.GetCPUCount(), sdl.GetPlatform())
		window, err := sdl.CreateWindow(GameName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,800, 600, sdl.WINDOW_SHOWN)
		if err != nil {
			log.Panicln(err)
		}
		defer window.Destroy()
		
		renderer, err := sdl.CreateRenderer(window, -1, 0)
		if err != nil {
			log.Panicln(err)
		}
		renderer.Clear()
		defer renderer.Destroy()
		var rendInfo sdl.RendererInfo
		err = renderer.GetRendererInfo(&rendInfo)
		if err != nil {
			log.Panicln(err)
		}
		log.Printf("Rendering with \"%s\". Max Texture Size %dx%d\n", rendInfo.Name, rendInfo.MaxTextureWidth, rendInfo.MaxTextureWidth)
		
		GameCam = NewCamera()
		InitProvences()
		InitNations()
		game_loop(renderer)
}