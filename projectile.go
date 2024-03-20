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
	dx := 0
	dy := 0

	switch p.Direction {
	case Up:
		dy = -1
		if p.Pixels[0].Y < BoxTop {
			return
		}
	case Down:
		dy = 1
		if p.Pixels[0].Y > BoxBottom {
			return
		}
	case Left:
		dx = -1
		if p.Pixels[0].X < BoxLeft {
			return
		}
	case Right:
		dx = 1
		if p.Pixels[1].X > BoxRight {
			return
		}
	}

	for _, pixel := range p.Pixels {
		box.Screen.SetContent(pixel.X, pixel.Y, ' ', nil, box.Style)

		pixel.X += dx
		pixel.Y += dy
	}
}
