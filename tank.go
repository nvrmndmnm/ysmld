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
	tankStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)
	dx := 0
	dy := 0
	for _, pixel := range t.Pixels {

		switch t.Direction {
		case Up:
			dy = -1
			box.Screen.SetContent(pixel.X, pixel.Y+1, ' ', nil, box.Style)
			box.Screen.SetContent(pixel.X+dx, pixel.Y+dy, '\u2588', nil, tankStyle)
		case Down:
			dy = 1
			box.Screen.SetContent(pixel.X, pixel.Y-1, ' ', nil, box.Style)
			box.Screen.SetContent(pixel.X+dx, pixel.Y+dy, '\u2588', nil, tankStyle)
		case Left:
			dx = -1
			box.Screen.SetContent(pixel.X+1, pixel.Y, ' ', nil, box.Style)
			box.Screen.SetContent(pixel.X+dx, pixel.Y+dy, '\u2588', nil, tankStyle)

		case Right:
			dx = 1
			box.Screen.SetContent(pixel.X-1, pixel.Y, ' ', nil, box.Style)
			box.Screen.SetContent(pixel.X+dx, pixel.Y+dy, '\u2588', nil, tankStyle)

		}
		pixel.X += dx
		pixel.Y += dy
		// if pixel.X <= BoxLeft || pixel.X >= BoxRight-5 || pixel.Y <= BoxTop || pixel.Y+2 >= BoxBottom {
		// break
		// }
	}

	// t.moveTurret(t.Pixels[0].X, t.Pixels[0].Y, t.Pixels[0].Style)
}

func (t *Tank) Shoot() *Projectile {
	if t.ShotsFired < t.Cooldown {
		return nil
	}
	t.ShotsFired = 0

	var x, y int

	switch t.Direction {
	case Up:
		x = t.Pixels[len(t.Pixels)-2].X
		y = t.Pixels[len(t.Pixels)-2].Y - 1
	case Down:
		x = t.Pixels[len(t.Pixels)-1].X
		y = t.Pixels[len(t.Pixels)-1].Y + 1
	case Left:
		x = t.Pixels[len(t.Pixels)-2].X - 1
		y = t.Pixels[len(t.Pixels)-2].Y
	case Right:
		x = t.Pixels[len(t.Pixels)-1].X + 1
		y = t.Pixels[len(t.Pixels)-1].Y
	}

	return NewProjectile(x, y, t.Direction)
}
