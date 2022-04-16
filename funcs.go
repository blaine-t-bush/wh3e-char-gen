package main

import (
	"math/rand"
)

// Roll <n>d<f>, i.e. generate <n> random integers from 1 to <f> inclusive and
// sum them.
func D(n int, f int) int {
	var total int
	for i := 0; i < n; i++ {
		total = total + rand.Intn(f+1)
	}

	return total
}

// Randomly generate six attribute scores with 3d6.
func GenerateAttributeScores() map[string]int {
	return map[string]int{
		"str": D(3, 6),
		"dex": D(3, 6),
		"con": D(3, 6),
		"int": D(3, 6),
		"wis": D(3, 6),
		"cha": D(3, 6),
	}
}

// Randomly generate hitpoints. The hitpoints are generated by rolling a number
// of d6 equal to hitDice, and adding bonusHitPoints. The minimum returned value
// is equal to previousHitPoints+1. At level 1, hit points are maximized.
func GenerateHitPoints(hitDice int, bonusHitPoints int, level int, previousHitPoints int) int {
	if level == 1 {
		return 6*hitDice + bonusHitPoints
	} else {
		newHitPoints := D(hitDice, 6) + bonusHitPoints
		if newHitPoints <= previousHitPoints {
			return previousHitPoints + 1
		} else {
			return newHitPoints
		}
	}
}
