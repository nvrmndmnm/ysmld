package main

import "github.com/gdamore/tcell/v2"

type Dot struct {
	X     int
	Y     int
	Style tcell.Style
	PrevX int
	PrevY int
}

func (d *Dot) Draw(s tcell.Screen) {
	d.PrevX = d.X
	d.PrevY = d.Y
	s.SetContent(d.X, d.Y, '*', nil, d.Style)
}

func (d *Dot) Clear(s tcell.Screen, boxStyle tcell.Style) {
	s.SetContent(d.PrevX, d.PrevY, ' ', nil, boxStyle)
}
