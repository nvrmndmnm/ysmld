package main

import "github.com/gdamore/tcell/v2"

type Dot struct {
	X int
	Y int
	Style tcell.Style
}

func (d *Dot) Draw(s tcell.Screen) {
	s.SetContent(d.X, d.Y, '*', nil, d.Style)
}
