package main

import (
	"math/rand"
)

func GenerateName() string {
	return "Placeholder"
}

func GenerateClass() Class {
	classNames := [3]string{
		"deft",
		"strong",
		"wise",
	}

	return classes[classNames[rand.Intn(len(classNames))]]
}

func GenerateSpecies() string {
	// Normalize probabilities.
	var totalProbability float64 = 0
	for _, species := range speciess {
		totalProbability += species.probability
	}

	// Randomly select a species.
	var cumulativeProbability float64 = 0
	var name string
	roll := rand.Float64()
	for _, species := range speciess {
		cumulativeProbability += species.probability / totalProbability
		if roll <= cumulativeProbability {
			name = species.name
			break
		}
	}

	return name
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
func GenerateHitPoints(hitDice int, bonusHitPoints int, level int, className string, conScore int, previousHitPoints int) int {
	if className == "Strong" && conScore >= 16 {
		bonusHitPoints += 2
	} else if className == "Strong" && conScore >= 13 {
		bonusHitPoints += 1
	}

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

func GenerateVocation() string {
	return vocations_[rand.Intn(len(vocations_)-1)]
}

func GenerateAffiliation(existingAffiliations []string) string {
	return GetUniqueElement(affiliations_, existingAffiliations)
}

func SelectGroupAttribute(groups []Group, excludes []string) string {
	// Determine which attributes have fewer than two groups already.
	groupAttributeCounts := map[string]int{
		"Strength":     0,
		"Dexterity":    0,
		"Constitution": 0,
		"Intelligence": 0,
		"Wisdom":       0,
		"Charisma":     0,
	}

	for _, group := range groups {
		groupAttributeCounts[group.attribute]++
	}

	var availableAttributes []string

	for attribute, count := range groupAttributeCounts {
		if count < 2 {
			availableAttributes = append(availableAttributes, attribute)
		}
	}

	// Exclude any attributes if that argument has been passed.
	if len(excludes) > 0 {
		var modifiedAvailableAttributes []string
		var include bool
		for _, attr := range availableAttributes {
			include = true
			for _, exclude := range excludes {
				if attr == exclude {
					include = false
					break
				}
			}

			if include {
				modifiedAvailableAttributes = append(modifiedAvailableAttributes, attr)
			}
		}

		availableAttributes = modifiedAvailableAttributes
	}

	// Randomly select one of those attributes.
	return availableAttributes[rand.Intn(len(availableAttributes)-1)]
}

func GenerateLanguages(intScore int) []string {
	languageCount := 1
	if intScore >= 16 {
		languageCount += 2
	} else if intScore >= 13 {
		languageCount += 1
	}

	languages := []string{"Low Imperial"} // Everyone knows the common tongue.
	for i := 1; i < languageCount; i++ {
		languages = append(languages, GetUniqueElement(languages_, languages))
	}

	return languages
}