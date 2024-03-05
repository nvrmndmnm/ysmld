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

func (p *Projectile) Move() {
	switch p.Direction {
	case Up:
		for _, pixel := range p.Pixels {
			if pixel.Y > BoxTop {
				pixel.Y--
			} else {
				p.Pixels = nil
			}
		}
	case Down:
		for _, pixel := range p.Pixels {
			if pixel.Y < BoxBottom {
				pixel.Y++
			} else {
				p.Pixels = nil
			}
		}
	case Left:
		for _, pixel := range p.Pixels {
			if pixel.X > BoxLeft {
				pixel.X -= 2
			} else {
				p.Pixels = nil
			}
		}
	case Right:
		for _, pixel := range p.Pixels {
			if pixel.X < BoxRight {
				pixel.X += 2
			} else {
				p.Pixels = nil
			}
		}
	}
}
