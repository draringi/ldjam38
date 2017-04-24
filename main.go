package main

import (
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
	"log"
	"github.com/kardianos/osext"
)

const (
	GameName = "War Never Changes"
)

var GameCam *Camera

var (
	execPath string
	gameWindow *sdl.Window
)

func main() {
		var err error
		execPath, err = osext.ExecutableFolder()
		if err != nil {
			log.Println(err)
		}
		err = sdl.Init(sdl.INIT_VIDEO|sdl.INIT_TIMER|sdl.INIT_AUDIO)
		if err != nil {
			log.Panicf("Failed to initialize SDL: %v\n", err)
		}
		defer sdl.Quit()
		err = ttf.Init()
		if err != nil {
			log.Panicf("Unable to initialize font engine: %v\n", err)
		}
		defer ttf.Quit()
		var sdlVersion sdl.Version
		sdl.GetVersion(&sdlVersion)
		log.Printf("SDL Initialized. Version: %d.%d.%d\n", sdlVersion.Major, sdlVersion.Minor, sdlVersion.Patch)
		log.Printf("CPU Count: %d, OS: %s\n", sdl.GetCPUCount(), sdl.GetPlatform())
		gameWindow, err = sdl.CreateWindow(GameName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,WindowWidth, WindowHeight, sdl.WINDOW_SHOWN)
		if err != nil {
			log.Panicln(err)
		}
		defer gameWindow.Destroy()
		
		renderer, err := sdl.CreateRenderer(gameWindow, -1, 0)
		if err != nil {
			log.Panicln(err)
		}
		renderer.Clear()
		defer renderer.Destroy()
		InitButtons(renderer)
		LoadSymbols(renderer)
		var rendInfo sdl.RendererInfo
		err = renderer.GetRendererInfo(&rendInfo)
		if err != nil {
			log.Panicln(err)
		}
		log.Printf("Rendering with \"%s\". Max Texture Size %dx%d\n", rendInfo.Name, rendInfo.MaxTextureWidth, rendInfo.MaxTextureWidth)
		
		game_loop(renderer)
		CleanupSymbols()
		CleanupButtons()
}