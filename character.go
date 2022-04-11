package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomClass() Class {
	classNames := [3]string{
		"deft",
		"strong",
		"wise",
	}

	return classes[classNames[rand.Intn(len(classNames))]]
}

func main() {
	// Seed the randomizer.
	rand.Seed(time.Now().UnixNano())
	// Randomly select a class.
	class := randomClass()
	// Randomly select a level.
	level := rand.Intn(10) + 1
	levelInfo := class.levels[level]

	fmt.Println("Class:", class.name)
	fmt.Println("Level:", level)
	fmt.Println(levelInfo)
}
