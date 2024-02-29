package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	BoxLeft   = 10
	BoxTop    = 1
	BoxRight  = 42
	BoxBottom = 20
)

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, ' ', nil, style)
		s.SetContent(col, y2, ' ', nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, ' ', nil, style)
		s.SetContent(x2, row, ' ', nil, style)
	}

	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, ' ', nil, style)
		s.SetContent(x2, y1, ' ', nil, style)
		s.SetContent(x1, y2, ' ', nil, style)
		s.SetContent(x2, y2, ' ', nil, style)
	}
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	drawBox(s, BoxLeft, BoxTop, BoxRight, BoxBottom, boxStyle)

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	tank := NewTank(11, 17)
	tank.Draw(s)
	var projectile *Projectile

	gameObject := &Object{
		Pixels: []*Pixel{
			{X: 15, Y: 15, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
			// {X: 16, Y: 15, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
			// {X: 15, Y: 16, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
			// {X: 16, Y: 16, Style: tcell.StyleDefault.Foreground(tcell.ColorRed)},
		},
	}

	gameObject.Draw(s)

	quitCh := make(chan struct{})

	go func() {
		for {
			s.Show()

			ev := s.PollEvent()

			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				dx, dy := 0, 0

				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(quitCh)
					return
				} else if ev.Key() == tcell.KeyCtrlL {
					s.Sync()
				} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
					s.Clear()
				} else if ev.Rune() == 'H' || ev.Rune() == 'h' {
					// move left
					canMove := true
					for _, pixel := range tank.Pixels {
						if pixel.X-1 <= BoxLeft {
							canMove = false
							break
						}
					}
					if canMove {
						dx = -1
					}
				} else if ev.Rune() == 'J' || ev.Rune() == 'j' {
					// move down
					canMove := true
					for _, pixel := range tank.Pixels {
						if pixel.Y+1 >= BoxBottom {
							canMove = false
							break
						}
					}
					if canMove {
						dy = 1
					}
				} else if ev.Rune() == 'K' || ev.Rune() == 'k' {
					// move up
					canMove := true
					for _, pixel := range tank.Pixels {
						if pixel.Y-1 <= BoxTop {
							canMove = false
							break
						}
					}
					if canMove {
						dy = -1
					}
				} else if ev.Rune() == 'L' || ev.Rune() == 'l' {
					// move right
					canMove := true
					for _, pixel := range tank.Pixels {
						if pixel.X+1 >= BoxRight {
							canMove = false
							break
						}
					}
					if canMove {
						dx = 1
					}
				} else if ev.Rune() == ' ' {
					// shoot
					projectile = tank.Shoot()
				}

				tank.ClearPrevious(s, boxStyle, dx, dy)
				tank.Move(dx, dy)
				tank.Draw(s)

				s.Show()
			}

		}
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-quitCh:
			ticker.Stop()
			return
		case <-ticker.C:
			if projectile != nil {
				projectile.ClearPrevious(s, boxStyle, 0, 1)
				projectile.Move()
				projectile.Draw(s)
			}
			tank.Draw(s)
			s.Show()
		}
	}
}
