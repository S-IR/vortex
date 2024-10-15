package vortex

import (
	"image/color"

	"golang.org/x/image/font"
)

type Styles struct {
	//Margin
	Margin, Mt, Mb, Ml, Mr int
	//Padding
	Padding, Pt, Pb, Pl, Pr int

	//Width, Height
	W, H int
	//X, Y
	X, Y  int
	Bg    color.RGBA
	Color color.RGBA

	Font string
	Text int

	fontFaceCache font.Face
	Transitions   Transitions
}

type StyleField int

const (
	Margin StyleField = iota
	Mt
	Mb
	Ml
	Mr
	Padding
	Pt
	Pb
	Pl
	Pr
	W
	H
	X
	Y
	Bg
	Color
	Font
	Text
	All
)

type TimingFunction int

const (
	Linear TimingFunction = iota
)

type TransitionProperty struct {
	TimingFunction TimingFunction
	Delay          int //in ms
	Duration       int //in ms
}
type Transitions = map[StyleField]TransitionProperty
