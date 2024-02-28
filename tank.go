// tank.go
package main

import "github.com/gdamore/tcell/v2"

type Tank struct {
	Object
}

func NewTank(x, y int) *Tank {
	tank := &Tank{}

	tankStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)
	turretStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkKhaki)

	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 6; dx++ {
			tank.Pixels = append(tank.Pixels, &Pixel{X: x + dx, Y: y + dy, Style: tankStyle})
		}
	}

	tank.Pixels = append(tank.Pixels,
		&Pixel{X: x + 2, Y: y - 1, Style: turretStyle},
		&Pixel{X: x + 3, Y: y - 1, Style: turretStyle})

	return tank
}

func (t *Tank) Move(dx, dy int) {
	for _, pixel := range t.Pixels {
		pixel.X += dx
		pixel.Y += dy
	}
}
