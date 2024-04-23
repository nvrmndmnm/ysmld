package main

import (
	"math/rand"
	"time"
)

var npcTanks []*Tank

var spawnPoints = [][2]int{
	{BoxLeft + 1, BoxTop + 1},
	{BoxRight - 6, BoxTop + 1},
	{BoxLeft + 1, BoxBottom - 3},
	{BoxRight - 6, BoxBottom - 3},
}

func getSpawnPoint() (int, int, bool) {
	if len(spawnPoints) == 0 {
		return 0, 0, false
	}

	randomIndex := rand.Intn(len(spawnPoints))
	spawnPoint := spawnPoints[randomIndex]
	spawnPoints = append(spawnPoints[:randomIndex], spawnPoints[randomIndex+1:]...)

	return spawnPoint[0], spawnPoint[1], true
}

func spawn(box *Box) {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		if len(npcTanks) < 4 {
			x, y, ok := getSpawnPoint()
			if ok {
				npcTank := NewTank(x, y)
				npcTank.Draw(box)
				npcTanks = append(npcTanks, npcTank)
			}
		}
	}
}

func despawn(tank *Tank, box *Box) {
	for i, npcTank := range npcTanks {
		if npcTank == tank {
			spawnPoints = append(spawnPoints, [2]int{tank.Pixels[0].X, tank.Pixels[0].Y})
			npcTanks = append(npcTanks[:i], npcTanks[i+1:]...)
			tank.Clear(box)

			break
		}
	}
}

func isHit(tank *Tank, ammo *Projectile) bool {
	for _, tankPixel := range tank.Pixels {
		for _, ammoPixel := range ammo.Pixels {
			if tankPixel.X == ammoPixel.X && tankPixel.Y == ammoPixel.Y {
				return true
			}
		}
	}
	return false
}
