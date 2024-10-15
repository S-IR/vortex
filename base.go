package vortex

import (
	"image/color"
	"math"
	"time"
)

type StylesModification = func(start Styles) Styles

// Component is a generic struct that can be used for both Div and P
type base[BodyType any, Self any] struct {
	Body    BodyType
	OnClick func(component Self, e *MouseEvent)
	OnHover func(component Self, e *MouseEvent)

	Styles      Styles
	StylesHover StylesModification

	hovered           bool
	hoveredFlipTime   time.Time
	hoverStylesCached Styles

	curX, curY int
}
type getCurrentStyleInput struct {
	Styles      Styles
	StylesHover StylesModification

	Hovered         bool
	HoveredFlipTime time.Time
}

func getCurrentStyle(input getCurrentStyleInput) Styles {
	if input.StylesHover == nil {
		return input.Styles
	}
	hoverStyles := input.StylesHover(input.Styles)

	if input.HoveredFlipTime.IsZero() {
		if input.Hovered {
			return hoverStyles
		}
		return input.Styles
	}

	elapsed := time.Since(input.HoveredFlipTime)

	result := input.Styles
	allComplete := true
	for field, value := range input.Styles.Transitions {
		interpolationFactor := 1.0
		durationInNanoseconds := time.Duration(value.Duration) * time.Millisecond
		if elapsed < durationInNanoseconds {
			interpolationFactor = math.Min(1, float64(elapsed)/float64(durationInNanoseconds))
			if !input.Hovered {
				interpolationFactor = 1 - interpolationFactor
			}
			allComplete = false
		} else if !input.Hovered {
			interpolationFactor = 0
		}

		interpolateStyle(&result, field, hoverStyles, interpolationFactor)
	}

	if allComplete {
		if input.Hovered {
			return hoverStyles
		}
		return input.Styles
	}

	return result
}

func interpolateStyle(s *Styles, field StyleField, hoverStyles Styles, t float64) {
	switch field {
	case W:
		s.W = interpolate(s.W, hoverStyles.W, t)
	case H:
		s.H = interpolate(s.H, hoverStyles.H, t)
	case Bg:
		s.Bg = interpolateColor(s.Bg, hoverStyles.Bg, t)
	case Color:
		s.Color = interpolateColor(s.Color, hoverStyles.Color, t)
	case All:
		s.W = interpolate(s.W, hoverStyles.W, t)
		s.H = interpolate(s.H, hoverStyles.H, t)
		s.Bg = interpolateColor(s.Bg, hoverStyles.Bg, t)
		s.Color = interpolateColor(s.Color, hoverStyles.Color, t)
	}
}

func interpolate(start, end int, t float64) int {
	if t <= 0 {
		return start
	}
	if t >= 1 {
		return end
	}
	return int(float64(start) + (float64(end)-float64(start))*t)
}

// Interpolates between two color.RGBA values
func interpolateColor(start, end color.RGBA, t float64) color.RGBA {
	t = clamp(t, 0, 1) // Ensure t is within bounds
	return color.RGBA{
		R: uint8(float64(start.R) + (float64(end.R)-float64(start.R))*t),
		G: uint8(float64(start.G) + (float64(end.G)-float64(start.G))*t),
		B: uint8(float64(start.B) + (float64(end.B)-float64(start.B))*t),
		A: uint8(float64(start.A) + (float64(end.A)-float64(start.A))*t),
	}
}

// Clamp value between min and max
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
