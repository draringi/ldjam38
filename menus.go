package main

import (
	"path"
	"log"
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
	"fmt"
)

const (
	DroidFont = "res/DroidSansMono.ttf"
	ButtonWidth = 100
	ButtonHeight = 50
	ButtonCentreX = ButtonWidth/2
	ButtonCentreY = ButtonHeight/2
	MenuWidth = 300
	MenuHeight = 400
)

type ButtonAction func(state *GameState)

var (
	buttons map[ButtonID]*sdl.Texture
	buttonActions map[ButtonID]ButtonAction
	buttonBindings []*ButtonBinding
	provenceButtonBindings []*ButtonBinding
	DroidSansMono *ttf.Font
)

type ButtonID int

const (
	ButtonQuit ButtonID = iota
	ButtonEndTurn
	ButtonResume
	ButtonNewGame
	ButtonHelp
	ButtonTrain
	ButtonBuild
	ButtonArmies
)

const (
	ButtonMenuX = WindowWidth/2 - ButtonWidth/2
	ButtonMenuY = WindowHeight/2-MenuHeight/2
	ButtonQuitX = ButtonMenuX
	ButtonQuitY = ButtonNewGameY + 2 * ButtonHeight
	ButtonEndTurnX = WindowWidth - ButtonWidth
	ButtonEndTurnY = WindowHeight - ButtonHeight
	ButtonResumeX = ButtonMenuX
	ButtonResumeY = ButtonMenuY + int32(1.5 * ButtonHeight)
	ButtonNewGameX = ButtonMenuX
	ButtonNewGameY = ButtonResumeY + 2 * ButtonHeight
	ButtonHelpX = ButtonMenuX
	ButtonHelpY = ButtonNewGameY + 2 * ButtonHeight
)

func DummyAction(state *GameState) {
	log.Println("Dummy Button Called")
}

func NewGameAction(state *GameState) {
	*state = GameReset
}

func QuitAction(state *GameState) {
	*state = GameQuit
}

func ResumeAction(state *GameState) {
	*state = GamePlay
}

func NewButtonBinding(x, y int32, ID ButtonID, state GameState) *ButtonBinding {
	b := new(ButtonBinding)
	b.ID = ID
	b.RequiredState = state
	b.x1 = x
	b.x2 = x + ButtonWidth
	b.y1 = y
	b.y2 = y + ButtonHeight
	return b
}

func AddButtonBindings() {
	buttonBindings = []*ButtonBinding{}
	buttonBindings = append(buttonBindings, NewButtonBinding(ButtonQuitX, ButtonQuitY, ButtonQuit, GameMenu))
	buttonBindings = append(buttonBindings, NewButtonBinding(ButtonResumeX, ButtonResumeY, ButtonResume, GameMenu))
	buttonBindings = append(buttonBindings, NewButtonBinding(ButtonNewGameX, ButtonNewGameY, ButtonNewGame, GameMenu))
	buttonBindings = append(buttonBindings, NewButtonBinding(ButtonEndTurnX, ButtonEndTurnY, ButtonEndTurn, GamePlay))
}

func CheckButtonPress (x, y int32, state *GameState) bool {
	for _, button := range(buttonBindings) {
		if *state == button.RequiredState && x > button.x1 && x < button.x2 && y > button.y1 && y < button.y2 {
			action := buttonActions[button.ID]
			if action != nil {
				action(state)
			}
			return true
		}
	}
	return false
}


func NewButtonSurface() *sdl.Surface {
	surface, err := sdl.CreateRGBSurface(0, ButtonWidth, ButtonHeight, 32, 0, 0, 0, 0xff000000)
	if err != nil {
		log.Panicln(err)
	}
	rect := sdl.Rect{0, 0, ButtonWidth, ButtonHeight}
	surface.FillRect(&rect, colourButton.Uint32())
	return surface
}

func MakeButton(r *sdl.Renderer, str string, ID ButtonID, action ButtonAction) {
	surface := NewButtonSurface()
	text, err := DroidSansMono.RenderUTF8_Blended(str, colourWhite)
	if err != nil {
		log.Panicln(err)
	}
	x := ButtonCentreX - text.W/2
	y := ButtonCentreY - text.H/2
	
	textLoc := sdl.Rect{x, y, text.W, text.H}
	textRect := sdl.Rect{0,0,text.W, text.H}
	text.Blit(&textRect, surface, &textLoc)
	text.Free()
	button, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		log.Panicln(err)
	}
	surface.Free()
	buttons[ID] = button
	buttonActions[ID] = action
}

func InitButtons(r *sdl.Renderer) {
	var err error
	buttons = make(map[ButtonID]*sdl.Texture)
	buttonActions = make(map[ButtonID]ButtonAction)
	fontPath := path.Join(execPath, DroidFont)
	DroidSansMono, err = ttf.OpenFont(fontPath, 16)
	if err != nil {
		log.Panicln(err)
	}
	MakeButton(r, "Quit", ButtonQuit, QuitAction)
	MakeButton(r, "End Turn", ButtonEndTurn, EndOfTurnAction)
	MakeButton(r, "Resume", ButtonResume, ResumeAction)
	MakeButton(r, "New Game", ButtonNewGame, NewGameAction)
	MakeButton(r, "Help", ButtonHelp, DummyAction)
	MakeButton(r, "Build", ButtonBuild, DummyAction)
	MakeButton(r, "Train", ButtonTrain, DummyAction)
	MakeButton(r, "Armies", ButtonArmies, ArmiesAction)
	
	AddButtonBindings()
	AddProvenceButtonBindings()
}

func CleanupButtons() {
	for _, button := range(buttons) {
		button.Destroy()
	}
	buttons = nil
	buttonActions = nil
	provenceButtonBindings = nil
	buttonBindings = nil
	DroidSansMono.Close()
}

type ButtonBinding struct {
	x1 int32
	x2 int32
	y1 int32
	y2 int32
	ID ButtonID
	RequiredState GameState
}

const (
	ButtonArmiesX = 0
	ButtonArmiesY = WindowHeight - ButtonHeight
	ButtonBuildX = ButtonWidth + 10
	ButtonBuildY = WindowHeight - ButtonHeight
	ButtonTrainX = 2 * ButtonWidth + 20
	ButtonTrainY = WindowHeight - ButtonHeight
)

func AddProvenceButtonBindings() {
	provenceButtonBindings = []*ButtonBinding{}
	provenceButtonBindings = append(provenceButtonBindings, NewButtonBinding(ButtonArmiesX, ButtonArmiesY, ButtonArmies, GamePlay))
	provenceButtonBindings = append(provenceButtonBindings, NewButtonBinding(ButtonBuildX, ButtonBuildY, ButtonBuild, GamePlay))
	provenceButtonBindings = append(provenceButtonBindings, NewButtonBinding(ButtonTrainX, ButtonTrainY, ButtonTrain, GamePlay))
}

func CheckProvenceMenu(x, y int32, state *GameState) bool {
	if !(x < ProvenceMenuWidth && y > (WindowHeight - ProvenceMenuHeight)) {
		return false
	}
	player := Nations["PLA"]
	if player == nil || SelectedProvence.Owner != player {
		return true
	}
	for _, button := range(provenceButtonBindings) {
		if x > button.x1 && x < button.x2 && y > button.y1 && y < button.y2 {
			action := buttonActions[button.ID]
			if action != nil {
				action(state)
			}
			return true
		}
	}
	return true
}

func ArmiesAction(state *GameState) {
	player := Nations["PLA"]
	if SelectedProvence == nil || player == nil {
		return
	}
	friendlies := FilterArmiesPresent(player, SelectedProvence)
	msgData := new(sdl.MessageBoxData)
	msgData.NumButtons = int32(len(friendlies) + 1)
	msgData.Window = gameWindow
	msgData.Title = fmt.Sprintf("Army Selection: %s", SelectedProvence.Name)
	msgData.Message = "Please select an army"
	msgData.Buttons = make([]sdl.MessageBoxButtonData, len(friendlies) + 1)
	for i, army := range(friendlies){
		var buttonData sdl.MessageBoxButtonData
		buttonData.ButtonId = int32(i)
		buttonData.Text = army.Name
		msgData.Buttons[i] = buttonData
	}
	msgData.Buttons[len(friendlies)] = sdl.MessageBoxButtonData{sdl.MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT, int32(len(friendlies) + 1), "Cancel"}
	err, id := sdl.ShowMessageBox(msgData)
	if err != nil {
		log.Panicln(err)
	}
	if id >= int32(len(friendlies)) {
		return
	}
	SelectedArmy = friendlies[int(id)]
	log.Printf("Selected army \"%s\"\n", SelectedArmy.Name)
}

func EndOfTurnAction(state *GameState) {
	player := Nations["PLA"]
	if player == nil {
		log.Println("End of whose turn?")
		return
	}
	player.EndTurn()
	*state = GameEoT
}