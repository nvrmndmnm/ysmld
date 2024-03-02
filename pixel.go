package main

import (
	"github.com/gdamore/tcell/v2"
)

type Pixel struct {
	X     int
	Y     int
	Style tcell.Style
}

func (p *Pixel) Draw(s tcell.Screen, i int) {
	s.SetContent(p.X, p.Y, '\u2588', nil, p.Style)
}

type Object struct {
	Pixels []*Pixel
}

func (g *Object) Draw(s tcell.Screen) {
	for i, pixel := range g.Pixels {
		pixel.Draw(s, i)
	}
}

func (g *Object) ClearPrevious(s tcell.Screen, style tcell.Style, dx, dy int) {
	if dx > 0 {
		dx -= 1
	}
	if dx < 0 {
		dx += 1
	}
	if dy > 0 {
		dy -= 1
	}
	if dy < 0 {
		dy += 1
	}

	for _, pixel := range g.Pixels {
		s.SetContent(pixel.X-dx, pixel.Y-dy, ' ', nil, style)
	}
}

func (g *Object) CanMove(dx, dy int) bool {
	for _, pixel := range g.Pixels {
		destX := pixel.X + dx
		destY := pixel.Y + dy
		if destX <= BoxLeft || destX >= BoxRight || destY <= BoxTop || destY >= BoxBottom {
			return false
		}
	}
	return true
}