package main

type Direction int

const (
	Up = iota
	Down
	Left
	Right
)

const (
	BoxLeft   = 11
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

	go runGame(box, playerTank, projectiles, quitCh)

	go spawn(box)
	go fire()

	handleAmmo(box, playerTank, projectiles, quitCh)
}
