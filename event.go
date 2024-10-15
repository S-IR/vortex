package vortex

type MouseEventType int

const (
	Hover MouseEventType = iota
	Click
)

type MouseEvent struct {
	X, Y int
	Type MouseEventType
}

func newMouseEvent(x, y int) *MouseEvent {

	eventType := Hover

	return &MouseEvent{
		X:    x,
		Y:    y,
		Type: eventType,
	}
}

type Rectangle struct {
	X      int // X coordinate of the top-left corner
	Y      int // Y coordinate of the top-left corner
	Width  int // Width of the rectangle
	Height int // Height of the rectangle
}

// Point represents a point with X and Y coordinates
type Point struct {
	X int // X coordinate of the point
	Y int // Y coordinate of the point
}

func (r Rectangle) IntersectsPoint(p Point) bool {
	return p.X >= r.X && p.X <= r.X+r.Width && p.Y >= r.Y && p.Y <= r.Y+r.Height
}
