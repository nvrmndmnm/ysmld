package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Direction int

const (
	Up = iota
	Down
	Left
	Right
)

const (
	BoxLeft   = 10
	BoxTop    = 1
	BoxRight  = 42
	BoxBottom = 20

	ScoreboardLeft   = 50
	ScoreboardTop    = 1
	ScoreboardRight  = 72
	ScoreboardBottom = 20

	ProjectileSpeed = 300
	ShootCooldown   = 5
	MaxProjectiles  = 5
)

func main() {
	box := NewBox()
	box.DrawBox(BoxLeft, BoxTop, BoxRight, BoxBottom)
	box.DrawBox(ScoreboardLeft, ScoreboardTop, ScoreboardRight, ScoreboardBottom)
	box.DisplayText("test\ntest")

	quit := func() {
		maybePanic := recover()
		box.Screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	playerTank := NewTank(23, 9)
	playerTank.IsPlayer = true
	playerTank.Draw(box)

	projectiles := make(chan *Projectile, MaxProjectiles)

	quitCh := make(chan struct{})

	go func() {
		for {
			box.Screen.Show()

			ev := box.Screen.PollEvent()

			switch ev := ev.(type) {
			case *tcell.EventResize:
				box.Screen.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(quitCh)
					return
				} else if ev.Key() == tcell.KeyCtrlL {
					box.Screen.Sync()
				} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
					box.Screen.Clear()
				} else if ev.Rune() == 'H' || ev.Rune() == 'h' {
					playerTank.Direction = Left
					playerTank.Move(box)
				} else if ev.Rune() == 'J' || ev.Rune() == 'j' {
					playerTank.Direction = Down
					playerTank.Move(box)
				} else if ev.Rune() == 'K' || ev.Rune() == 'k' {
					playerTank.Direction = Up
					playerTank.Move(box)
				} else if ev.Rune() == 'L' || ev.Rune() == 'l' {
					playerTank.Direction = Right
					playerTank.Move(box)
				} else if ev.Rune() == ' ' {
					projectile := playerTank.Shoot()

					if projectile != nil {
						select {
						case projectiles <- projectile:
							playerTank.ShotsFired++
						default:
							projectile = nil
						}
					}
				}

				playerTank.Draw(box)
				box.Screen.Show()
			}

		}
	}()

	go spawn(box)

	var ammoRack []*Projectile
	ticker := time.NewTicker(ProjectileSpeed * time.Millisecond)
	for {
		select {
		case projectile := <-projectiles:
			ammoRack = append(ammoRack, projectile)
		case <-ticker.C:
			for i := 0; i < len(ammoRack) && i < MaxProjectiles; i++ {
				if ammoRack[i].Pixels[0].Y > BoxTop &&
					ammoRack[i].Pixels[0].Y < BoxBottom &&
					ammoRack[i].Pixels[0].X > BoxLeft+1 &&
					ammoRack[i].Pixels[1].X < BoxRight-1 {

					for _, npcTank := range npcTanks {
						if isHit(npcTank, ammoRack[i]) {
							npcTank.Clear(box)
							despawn(npcTank, box)
							break
						}
					}

					ammoRack[i].Move(box)
					ammoRack[i].Draw(box)
				} else {
					ammoRack[i].Clear(box)
					ammoRack = ammoRack[i:]
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
