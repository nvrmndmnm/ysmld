package main

import (
	"github.com/gdamore/tcell/v2"
	"time"
)

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
		dx = -2
		if p.Pixels[0].X < BoxLeft {
			return
		}
	case Right:
		dx = 2
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

func handleAmmo(box *Box, playerTank *Tank, projectiles chan *Projectile, quitCh chan struct{}) {
	var ammoRack []*Projectile
	ticker := time.NewTicker(ProjectileSpeed * time.Millisecond)
	for {
		select {
		case projectile := <-projectiles:
			ammoRack = append(ammoRack, projectile)
		case <-ticker.C:
			for i := len(ammoRack) - 1; i >= 0; i-- {
				if ammoRack[i].Pixels[0].Y > BoxTop &&
					ammoRack[i].Pixels[0].Y < BoxBottom &&
					ammoRack[i].Pixels[0].X > BoxLeft+1 &&
					ammoRack[i].Pixels[1].X < BoxRight-1 {

					for _, npcTank := range npcTanks {
						if isHit(npcTank, ammoRack[i]) {
							ammoRack[i].Clear(box)
							ammoRack = append(ammoRack[:i], ammoRack[i+1:]...)

							despawn(npcTank, box)
							break
						}
					}

					if i < len(ammoRack) {
						ammoRack[i].Move(box)
						ammoRack[i].Draw(box)
					}
				} else {
					ammoRack[i].Clear(box)
					ammoRack = append(ammoRack[:i], ammoRack[i+1:]...)
				}
			}

			playerTank.Draw(box)
			box.Screen.Show()

		case <-quitCh:
			ticker.Stop()
			return
		}
	}
}
