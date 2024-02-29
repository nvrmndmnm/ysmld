package main

import "github.com/gdamore/tcell/v2"

type Tank struct {
	Object
	Direction string
}

func NewTank(x, y int) *Tank {
	tank := &Tank{Direction: "up"}

	tankStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)
	turretStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkKhaki)

	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 6; dx++ {
			tank.Pixels = append(tank.Pixels, &Pixel{X: x + dx, Y: y + dy, Style: tankStyle})
		}
	}

	tank.moveTurret(x, y, turretStyle)

	return tank
}

func (t *Tank) moveTurret(x, y int, style tcell.Style) {
	if len(t.Pixels) > 18 {
		t.Pixels = t.Pixels[:18]
	}

	switch t.Direction {
	case "up":
		t.Pixels = append(t.Pixels,
			&Pixel{X: x + 2, Y: y - 1, Style: style},
			&Pixel{X: x + 3, Y: y - 1, Style: style})
	case "down":
		t.Pixels = append(t.Pixels,
			&Pixel{X: x + 2, Y: y + 3, Style: style},
			&Pixel{X: x + 3, Y: y + 3, Style: style})
	case "left":
		t.Pixels = append(t.Pixels,
			&Pixel{X: x - 1, Y: y + 1, Style: style},
			&Pixel{X: x - 2, Y: y + 1, Style: style})
	case "right":
		t.Pixels = append(t.Pixels,
			&Pixel{X: x + 6, Y: y + 1, Style: style},
			&Pixel{X: x + 7, Y: y + 1, Style: style})
	}
}

func (t *Tank) Move(dx, dy int) {
	for _, pixel := range t.Pixels {
		pixel.X += dx
		pixel.Y += dy
	}

	if dx > 0 {
		t.Direction = "right"
	} else if dx < 0 {
		t.Direction = "left"
	} else if dy > 0 {
		t.Direction = "down"
	} else if dy < 0 {
		t.Direction = "up"
	}

	t.moveTurret(t.Pixels[0].X, t.Pixels[0].Y, t.Pixels[0].Style)
}
