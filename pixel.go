package main

import "github.com/gdamore/tcell/v2"

type Pixel struct {
	X     int
	Y     int
	Style tcell.Style
}

func (p *Pixel) Draw(s tcell.Screen) {
	s.SetContent(p.X, p.Y, '\u2588', nil, p.Style)
	s.SetContent(p.X+1, p.Y, '\u2588', nil, p.Style)
}

type Object struct {
	Pixels []*Pixel
}

func (g *Object) Draw(s tcell.Screen) {
	for _, pixel := range g.Pixels {
		pixel.Draw(s)
	}
}

func (g *Object) ClearPrevious(s tcell.Screen, style tcell.Style, dx, dy int) {
    for _, pixel := range g.Pixels {
        s.SetContent(pixel.X-dx, pixel.Y-dy, ' ', nil, style)
    }
}