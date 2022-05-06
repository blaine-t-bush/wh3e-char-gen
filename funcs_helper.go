package main

import (
	"math/rand"
)

// Roll <n>d<f>, i.e. generate <n> random integers from 1 to <f> inclusive and
// sum them.
func D(n int, f int) int {
	var total int
	for i := 0; i < n; i++ {
		total = total + rand.Intn(f) + 1
	}

	return total
}

func GetUniqueElement(possibilities []string, excludes []string) string {
	var newElement string
	var keepSeeking bool

	for {
		keepSeeking = false
		newElement = possibilities[rand.Intn(len(possibilities)-1)]

		for _, element := range excludes {
			if newElement == element {
				keepSeeking = true
				break
			}
		}

		if !keepSeeking {
			break
		}
	}

	return newElement
}
