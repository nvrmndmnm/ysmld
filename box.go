package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Box struct {
	Screen tcell.Screen
	Style  tcell.Style
}

func NewBox() *Box {

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	box := &Box{
		Screen: s,
		Style:  boxStyle,
	}

	return box
}

func (box Box) DrawBox(x1, y1, x2, y2 int) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			box.Screen.SetContent(col, row, ' ', nil, box.Style)
		}
	}

	for col := x1; col <= x2; col++ {
		box.Screen.SetContent(col, y1, ' ', nil, box.Style)
		box.Screen.SetContent(col, y2, ' ', nil, box.Style)
	}
	for row := y1 + 1; row < y2; row++ {
		box.Screen.SetContent(x1, row, ' ', nil, box.Style)
		box.Screen.SetContent(x2, row, ' ', nil, box.Style)
	}

	if y1 != y2 && x1 != x2 {
		box.Screen.SetContent(x1, y1, ' ', nil, box.Style)
		box.Screen.SetContent(x2, y1, ' ', nil, box.Style)
		box.Screen.SetContent(x1, y2, ' ', nil, box.Style)
		box.Screen.SetContent(x2, y2, ' ', nil, box.Style)
	}
}

func (box Box) DisplayText(str string) {
	box.ClearScoreboard()

	column := ScoreboardLeft + 1
	row := ScoreboardTop
	for _, r := range str {
		box.Screen.SetContent(column, row, r, nil, box.Style)
		column++
		if r == '\n' {
			column = ScoreboardLeft + 1
			row++
		}
	}
}

func (box Box) ClearScoreboard() {
	for y := ScoreboardTop; y < ScoreboardBottom; y++ {
		for x := ScoreboardLeft; x < ScoreboardRight; x++ {
			box.Screen.SetContent(x, y, ' ', nil, box.Style)
		}
	}
}
