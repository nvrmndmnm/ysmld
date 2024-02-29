package main

import "github.com/gdamore/tcell/v2"

type Projectile struct {
	Object
	Direction string
}

func NewProjectile(x, y int, direction string) *Projectile {
	projectile := &Projectile{Direction: direction}
	
	projectileStyle := tcell.StyleDefault.Foreground(tcell.ColorRed)
	
	projectile.Pixels = append(projectile.Pixels,
		&Pixel{X: x, Y: y, Style: projectileStyle},
		&Pixel{X: x + 1, Y: y, Style: projectileStyle},
	)

	return projectile
}

func (p *Projectile) Move() {
	switch p.Direction {
	case "up":
		for _, pixel := range p.Pixels {
			pixel.Y--
		}
	case "down":
		for _, pixel := range p.Pixels {
			pixel.Y++
		}
	case "left":
		for _, pixel := range p.Pixels {
			pixel.X--
		}
	case "right":
		for _, pixel := range p.Pixels {
			pixel.X++
		}
	}
}
