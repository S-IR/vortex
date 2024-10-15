package vortex

import (
	"os"
	"time"
	"unicode"

	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type P base[string, *P]

// func (p *P) draw() {
// 	lib.Assert(gApp != nil, "Calling draw when global app is nil")

// 	if p.Styles.fontFaceCache == nil {
// 		var err error

// 		fontPath, err := findfont.Find(p.Styles.Font)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fontData, err := os.ReadFile(fontPath)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fnt, err := truetype.Parse(fontData)
// 		if err != nil {
// 			panic(err)
// 		}

// 		p.Styles.fontFaceCache = truetype.NewFace(fnt, &truetype.Options{
// 			Size:    float64(p.Styles.Text),
// 			DPI:     72,
// 			Hinting: font.HintingFull,
// 		})
// 		p.Styles.W, p.Styles.H = measureText(p.Styles.fontFaceCache, p.Body)
// 	}

// 	chosenStyles := getCurrentStyle(getCurrentStyleInput{
// 		Styles:          p.Styles,
// 		StylesHover:     p.StylesHover,
// 		Hovered:         p.hovered,
// 		HoveredFlipTime: p.hoveredFlipTime,
// 	})

// 	font := &font.Drawer{
// 		Dst:  gApp.screenBuffer.RGBA(),
// 		Src:  &image.Uniform{chosenStyles.Color},
// 		Face: chosenStyles.fontFaceCache,
// 		Dot:  fixed.P(gApp.curX, gApp.curY+chosenStyles.Text),
// 	}

// 	font.DrawString(p.Body)
// 	gApp.curY += chosenStyles.Text

// }

func (p *P) on(e *MouseEvent) {

	box := Rectangle{
		X:      p.curX,
		Y:      p.curY,
		Width:  p.Styles.W,
		Height: p.Styles.H,
	}
	mousePos := Point{
		X: e.X,
		Y: e.Y,
	}

	intersects := box.IntersectsPoint(mousePos)
	defer func() {
		mouseEventIsHover := e.Type == Hover && intersects
		if p.hovered != mouseEventIsHover {
			p.hovered = mouseEventIsHover
			p.hoveredFlipTime = time.Now()
		}
	}()

	if !intersects {
		return
	}

	if e.Type == Click && p.OnClick != nil {
		p.OnClick(p, e)
	}

	if e.Type == Hover && p.OnHover != nil {
		p.OnHover(p, e)
	}

}

func (p *P) loadFont() {
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
	p.Styles.W, p.Styles.H = measureText(p.Styles.fontFaceCache, p.Body)
}

func measureText(fontFace font.Face, text string) (int, int) {
	var totalWidth, totalHeight fixed.Int26_6

	for _, char := range text {
		if unicode.IsGraphic(char) {
			bbox, _, _ := fontFace.GlyphBounds(char)
			totalWidth += bbox.Max.X - bbox.Min.X
			totalHeight = max(totalHeight, bbox.Max.Y-bbox.Min.Y)
		}
	}

	return totalWidth.Ceil(), totalHeight.Ceil()
}
