package vortex

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"github.com/s-ir/vortex/lib"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

type App struct {
	//Default: 1280x720
	Width, Height int
	//Default: Hello Vortex
	Title        string
	window       screen.Window
	screenBuffer screen.Buffer

	Body []Component

	curX, curY int
}

// Global app
var gApp *App

type Component interface {
	Draw()
}
type Styles struct {
	//Width, Height
	W, H int
	//X, Y
	X, Y  int
	Bg    color.RGBA
	Color color.RGBA

	Font string
	Text int

	fontFaceCache font.Face
}
type Div struct {
	Styles Styles
	Body   []Component
}

type P struct {
	Styles Styles
	Body   string
}

func (p *P) Draw() {
	lib.Assert(gApp != nil, "Calling draw when global app is nil")

	if p.Styles.fontFaceCache == nil {
		var err error

		fontPath, err := findfont.Find(p.Styles.Font)
		if err != nil {
			panic(err)
		}

		fontData, err := os.ReadFile(fontPath)
		if err != nil {
			panic(err)
		}
		fnt, err := truetype.Parse(fontData)
		if err != nil {
			panic(err)
		}
		p.Styles.fontFaceCache = truetype.NewFace(fnt, &truetype.Options{
			Size:    float64(p.Styles.Text),
			DPI:     72,
			Hinting: font.HintingFull,
		})

	}
	font := &font.Drawer{
		Dst:  gApp.screenBuffer.RGBA(),
		Src:  &image.Uniform{p.Styles.Color},
		Face: p.Styles.fontFaceCache,
		Dot:  fixed.P(gApp.curX, gApp.curY+p.Styles.Text),
	}

	font.DrawString(p.Body)
	gApp.curY += p.Styles.Text

}
func (d *Div) Draw() {
	lib.Assert(gApp != nil, "Calling draw when global app is nil")

	rect := image.Rect(gApp.curX, gApp.curY, gApp.curX+d.Styles.W, gApp.curY+d.Styles.H)

	draw.Draw(
		gApp.screenBuffer.RGBA(),    // The destination image (buffer).
		rect,                        // The area to fill.
		&image.Uniform{d.Styles.Bg}, // The source (fill color).
		image.Point{},               // Starting point.
		draw.Src,                    // Draw operation mode.
	)
	// gApp.curX += d.Styles.W
	gApp.curY += d.Styles.H

}

func (a *App) Stop() {
	a.window.Release()
	a.screenBuffer.Release()
	gApp = nil
}
func (a *App) Start() error {
	gApp = a

	switch true {
	case a.Width == 0:
		a.Width = 1280
		fallthrough
	case a.Height == 0:
		a.Height = 720
		fallthrough
	case a.Title == "":
		a.Title = "Hello, Vortex!"
	}
	var err error

	driver.Main(func(s screen.Screen) {

		a.window, err = s.NewWindow(&screen.NewWindowOptions{
			Title:  a.Title,
			Width:  a.Width,
			Height: a.Height,
		})
		if err != nil {
			return
		}

		size0 := image.Point{a.Width, a.Height}
		a.screenBuffer, err = s.NewBuffer(size0)
		if err != nil {
			return
		}

		draw.Draw(a.screenBuffer.RGBA(), a.screenBuffer.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

		for _, component := range a.Body {
			component.Draw()
		}
		// var sz size.Event

		for {
			e := a.window.NextEvent()

			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				a.window.Upload(image.Point{}, a.screenBuffer, a.screenBuffer.Bounds())
				a.window.Publish()
			// case size.Event:
			// sz = e
			case error:
				panic(e)

			}
		}

	})
	return err
}

func main() {
	fmt.Println("hello world")
}
