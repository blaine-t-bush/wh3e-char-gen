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
	// Start at level 1.
	level := 2
	levelInfo := class.levels[level]
	// Randomly generate attributes.
	attributeScores := GenerateAttributeScores()
	// Generate hit points.
	hitPoints := GenerateHitPoints(levelInfo.hitDie, levelInfo.bonusHitPoints, level, 0)

	fmt.Println("Class: ", class.name)
	fmt.Println("Level: ", level)
	fmt.Println("HP:    ", hitPoints)
	fmt.Println("AV:    ", levelInfo.attackValue)
	fmt.Println("SV:    ", levelInfo.savingThrow)
	fmt.Println("Slots: ", levelInfo.slotCount)
	fmt.Println("Groups:", levelInfo.groupCount)
	fmt.Println("Raises:", levelInfo.raiseCount)
	fmt.Println("# Attributes")
	fmt.Println(" - STR", attributeScores["str"])
	fmt.Println(" - DEX", attributeScores["dex"])
	fmt.Println(" - CON", attributeScores["con"])
	fmt.Println(" - INT", attributeScores["int"])
	fmt.Println(" - WIS", attributeScores["wis"])
	fmt.Println(" - CHA", attributeScores["cha"])
}
