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

	gameObject := &Object{
		Pixels: []*Pixel{
			{X: 15, Y: 15, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
			// {X: 16, Y: 15, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
			// {X: 15, Y: 16, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
			// {X: 16, Y: 16, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
		},
	}

	gameObject.Draw(box.Screen)

	quitCh := make(chan struct{})

	go func() {
		for {
			box.Screen.Show()

			ev := box.Screen.PollEvent()

			switch ev := ev.(type) {
			case *tcell.EventResize:
				box.Screen.Sync()
			case *tcell.EventKey:
				dx, dy := 0, 0

				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(quitCh)
					return
				} else if ev.Key() == tcell.KeyCtrlL {
					box.Screen.Sync()
				} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
					box.Screen.Clear()
				} else if ev.Rune() == 'H' || ev.Rune() == 'h' {
					// move left
					tank.Direction = Left
					// if tank.CanMove(-1, 0) {
					dx = -1
					// }
				} else if ev.Rune() == 'J' || ev.Rune() == 'j' {
					// move down
					tank.Direction = Down
					// if tank.CanMove(0, 1) {
					dy = 1
					// }
				} else if ev.Rune() == 'K' || ev.Rune() == 'k' {
					// move up
					tank.Direction = Up
					// if tank.CanMove(0, -1) {
					dy = -1
					// }
				} else if ev.Rune() == 'L' || ev.Rune() == 'l' {
					// move right
					tank.Direction = Right
					// if tank.CanMove(1, 0) {
					dx = 1
					// }
				} else if ev.Rune() == ' ' {
					// shoot
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

				tank.ClearPrevious(box.Screen, box.Style)
				tank.Move(dx, dy)
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
			for len(projectiles) > 0 {
				projectile := <-projectiles
				projectile.ClearPrevious(box.Screen, box.Style)
				projectile.Move()
				projectile.Draw(box.Screen)

				// Re-enqueue the projectile
				projectiles <- projectile
			}
			tank.Draw(box.Screen)
			box.Screen.Show()
		}
	}
}
