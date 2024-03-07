package main

import (
	// "fmt"

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
	Pixels    []*Pixel
	Direction Direction
}

func (g *Object) Draw(s tcell.Screen) {
	for i, pixel := range g.Pixels {
		pixel.Draw(s, i)
	}
}
