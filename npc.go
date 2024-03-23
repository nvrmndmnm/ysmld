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
				npcTank.Draw(box.Screen)
				npcTanks = append(npcTanks, npcTank)
			}
		}
	}
}
