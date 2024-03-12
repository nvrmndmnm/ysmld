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

	tank := NewTank(11, 17)
	tank.Draw(box.Screen)
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
					tank.Direction = Left
					tank.Move(box)
				} else if ev.Rune() == 'J' || ev.Rune() == 'j' {
					tank.Direction = Down
					tank.Move(box)
				} else if ev.Rune() == 'K' || ev.Rune() == 'k' {
					tank.Direction = Up
					tank.Move(box)
				} else if ev.Rune() == 'L' || ev.Rune() == 'l' {
					tank.Direction = Right
					tank.Move(box)
				} else if ev.Rune() == ' ' {
					projectile := tank.Shoot()

					if projectile != nil {
						select {
						case projectiles <- projectile:
							tank.ShotsFired++
						default:
							projectile = nil
						}
					}
				}

				tank.Draw(box.Screen)
				box.Screen.Show()
			}

		}
	}()

	ticker := time.NewTicker(ProjectileSpeed * time.Millisecond)
	for {
		select {
		case <-quitCh:
			ticker.Stop()
			return
		case <-ticker.C:
			projectile := <-projectiles
			projectile.Move(box)
			projectiles <- projectile

			tank.Draw(box.Screen)
			box.Screen.Show()
		}
	}
}
