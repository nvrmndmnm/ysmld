package main

import (
	"github.com/gdamore/tcell/v2"
)

type Tank struct {
	Object
	ShotsFired int
	Cooldown   int
}

func NewTank(x, y int) *Tank {
	tank := &Tank{
		ShotsFired: 0,
		Cooldown:   ShootCooldown,
	}
	tank.Direction = Up
	tankStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)

	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 6; dx++ {
			tank.Pixels = append(tank.Pixels, &Pixel{X: x + dx, Y: y + dy, Style: tankStyle})
		}
	}

	// tank.moveTurret(x, y, tankStyle)

	return tank
}

func (t *Tank) moveTurret(x, y int, style tcell.Style) {
	if len(t.Pixels) > 18 {
		t.Pixels = t.Pixels[:18]
	}

	switch t.Direction {
	case Up:
		t.Pixels = append(t.Pixels,
			&Pixel{X: x + 2, Y: y - 1, Style: style},
			&Pixel{X: x + 3, Y: y - 1, Style: style},
		)
	case Down:
		t.Pixels = append(t.Pixels,
			&Pixel{X: x + 2, Y: y + 3, Style: style},
			&Pixel{X: x + 3, Y: y + 3, Style: style},
		)
	case Left:
		t.Pixels = append(t.Pixels,
			&Pixel{X: x - 1, Y: y + 1, Style: style},
			&Pixel{X: x - 2, Y: y + 1, Style: style},
		)
	case Right:
		t.Pixels = append(t.Pixels,
			&Pixel{X: x + 6, Y: y + 1, Style: style},
			&Pixel{X: x + 7, Y: y + 1, Style: style},
		)
	}
}

func (t *Tank) Move(box *Box) {
	dx := 0
	dy := 0

	switch t.Direction {
	case Up:
		dy = -1
		if t.Pixels[0].Y+dy <= BoxTop-1 {
			return
		}
	case Down:
		dy = 1
		if t.Pixels[0].Y+dy >= BoxBottom-1 {
			return
		}
	case Left:
		dx = -1
		if t.Pixels[0].X+dx <= BoxLeft {
			return
		}
	case Right:
		dx = 1
		if t.Pixels[0].X+dx >= BoxRight-5 {
			return
		}
	}

	for _, pixel := range t.Pixels {
		box.Screen.SetContent(pixel.X-dx, pixel.Y, ' ', nil, box.Style)

		pixel.X += dx
		pixel.Y += dy

	}

	// t.moveTurret(t.Pixels[0].X, t.Pixels[0].Y, t.Pixels[0].Style)
}

func (t *Tank) Shoot() *Projectile {
	if t.ShotsFired > t.Cooldown {
		return nil
	}
	t.ShotsFired = 0

	var x, y int

	switch t.Direction {
	case Up:
		x = t.Pixels[0].X + 2
		y = t.Pixels[0].Y - 1
	case Down:
		x = t.Pixels[0].X + 3
		y = t.Pixels[0].Y + 3
	case Left:
		x = t.Pixels[0].X
		y = t.Pixels[0].Y + 1
	case Right:
		x = t.Pixels[0].X + 6
		y = t.Pixels[0].Y + 1
	}

	return NewProjectile(x, y, t.Direction)
}
