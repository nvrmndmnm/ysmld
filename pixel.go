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
	Pixels    []*Pixel
	Direction Direction
}

func (g *Object) Draw(box *Box) {
	for i, pixel := range g.Pixels {
		pixel.Draw(box.Screen, i)
	}
}

func (g *Object) Clear(box *Box) {
	for _, pixel := range g.Pixels {
		box.Screen.SetContent(pixel.X, pixel.Y, ' ', nil, box.Style)
	}
}
