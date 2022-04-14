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
func GenerateAttributeScores() [6]int {
	return [6]int{D(3, 6), D(3, 6), D(3, 6), D(3, 6), D(3, 6), D(3, 6)}
}
