package main

import (
	"math/rand"
	"time"
)

var npcTanks []*Tank

func getSpawnPoints() [][2]int {
	return [][2]int{
		{BoxLeft + 1, BoxTop + 1},
		{BoxRight - 6, BoxTop + 1},
		{BoxLeft + 1, BoxBottom - 3},
		{BoxRight - 6, BoxBottom - 3},
	}
}

func spawn(box *Box) {
	ticker := time.NewTicker(1 * time.Second)
	spawnPoints := getSpawnPoints()

	for range ticker.C {
		if len(npcTanks) < 4 {
			spawnPoint := spawnPoints[rand.Intn(len(spawnPoints))]
			npcTank := NewTank(spawnPoint[0], spawnPoint[1])
			npcTank.Draw(box.Screen)
			npcTanks = append(npcTanks, npcTank)
		}
	}
}
