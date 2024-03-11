package main

import "github.com/gdamore/tcell/v2"

type Projectile struct {
	Object
}

func NewProjectile(x, y int, direction Direction) *Projectile {
	projectile := &Projectile{}
	projectile.Direction = direction

	projectileStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)

	x2 := x - 1
	if direction == Up {
		x2 = x + 1
	}

	projectile.Pixels = append(projectile.Pixels,
		&Pixel{X: x, Y: y, Style: projectileStyle},
		&Pixel{X: x2, Y: y, Style: projectileStyle},
	)

	return projectile
}

func (p *Projectile) Move(box *Box) {
	projectileStyle := tcell.StyleDefault.Foreground(tcell.ColorRebeccaPurple)
	dx := 0
	dy := 0

	for _, pixel := range p.Pixels {
		switch p.Direction {
		case Up:
			dy = -1
		case Down:
			dy = 1
		case Left:
			dx = -1
		case Right:
			dx = 1
		}

		box.Screen.SetContent(pixel.X-dx, pixel.Y-dy, ' ', nil, box.Style)
		box.Screen.SetContent(pixel.X+dx, pixel.Y+dy, '\u2588', nil, projectileStyle)

		pixel.X += dx
		pixel.Y += dy
	}
}
