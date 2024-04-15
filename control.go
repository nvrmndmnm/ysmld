package main

import "github.com/gdamore/tcell/v2"

func runGame(box *Box, playerTank *Tank, projectiles chan *Projectile, quitCh chan struct{}) {
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
}
