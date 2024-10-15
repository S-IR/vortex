package vortex

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	//Default: 1280x720
	Width, Height int
	//Default: Hello Vortex
	Title string
	//Default: 60
	TargetFPS int
	// window       screen.Window
	// screenBuffer screen.Buffer

	Body []Component

	curX, curY int
	fps        int
}

func FPS() float64 {
	return float64(time.Second) / float64(rl.GetFrameTime())
}

// Global app
var gApp *App

type Component interface {
	draw()
	on(e *MouseEvent)
}

func (a *App) Stop() {
	rl.CloseWindow()
	gApp = nil
}

var mouseMoved = false
var isMouseOutside = false

func (a *App) Start() error {
	gApp = a

	// Set default values
	if a.Width == 0 {
		a.Width = 1280
	}
	if a.Height == 0 {
		a.Height = 720
	}
	if a.TargetFPS == 0 {
		a.TargetFPS = 60
	}
	if a.Title == "" {
		a.Title = "Hello, Vortex!"
	}

	// Initialize Raylib window
	rl.InitWindow(int32(a.Width), int32(a.Height), a.Title)
	rl.SetTargetFPS(int32(a.TargetFPS))

	for !rl.WindowShouldClose() {
		// Update input and events
		a.curX = 0
		a.curY = 0
		handleMouse()
		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for _, component := range a.Body {
			component.draw()
		}

		rl.EndDrawing()

	}

	a.Stop()
	return nil
}
func handleMouse() {
	mousePos := rl.GetMousePosition()

	isMouseOutside = (mousePos.X < 0 || mousePos.X >= float32(gApp.Width) || mousePos.Y < 0 || mousePos.Y >= float32(gApp.Height))

	if isMouseOutside {
		return
	}
	if !mouseMoved && (mousePos.X != 0 || mousePos.Y != 0) {
		mouseMoved = true
	}

	if mouseMoved {
		for _, comp := range gApp.Body {
			comp.on(&MouseEvent{
				X:    int(mousePos.X),
				Y:    int(mousePos.Y),
				Type: Hover,
			})
		}
	}

	// Handle mouse events
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for _, comp := range gApp.Body {
			comp.on(&MouseEvent{
				X:    int(mousePos.X),
				Y:    int(mousePos.Y),
				Type: Click,
			})
		}
	}

}
