package vortex

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/s-ir/vortex/lib"
)

type Div base[[]Component, *Div]

func (div *Div) on(e *MouseEvent) {

	box := Rectangle{
		X:      div.curX,
		Y:      div.curY,
		Width:  div.Styles.W,
		Height: div.Styles.H,
	}
	mousePos := Point{
		X: e.X,
		Y: e.Y,
	}

	intersects := box.IntersectsPoint(mousePos)

	if div.hovered != intersects {
		div.hovered = intersects
		div.hoveredFlipTime = time.Now()
	}

	if intersects && e.Type == Hover && div.OnHover != nil {
		div.OnHover(div, e)
	}

	if intersects && e.Type == Click && div.OnClick != nil {
		div.OnClick(div, e)
	}
}

// var idk = 0

func (div *Div) draw() {
	lib.Assert(gApp != nil, "Calling draw when global app is nil")

	// idk++

	chosenStyles := getCurrentStyle(getCurrentStyleInput{
		Styles:          div.Styles,
		StylesHover:     div.StylesHover,
		Hovered:         div.hovered,
		HoveredFlipTime: div.hoveredFlipTime,
	})

	// Define rectangle dimensions for Raylib drawing
	rect := rl.Rectangle{
		X:      float32(gApp.curX),
		Y:      float32(gApp.curY),
		Width:  float32(chosenStyles.W),
		Height: float32(chosenStyles.H),
	}
	// Draw the rectangle with the chosen background color
	rl.DrawRectangleRec(rect, rl.Color{
		R: chosenStyles.Bg.R,
		G: chosenStyles.Bg.G,
		B: chosenStyles.Bg.B,
		A: chosenStyles.Bg.A,
	})

	div.curX = gApp.curX
	div.curY = gApp.curY
	// Update the global app's current Y position
	gApp.curY += chosenStyles.H

	// Draw child components inside the Div
	for _, comp := range div.Body {
		comp.draw()
	}
}
