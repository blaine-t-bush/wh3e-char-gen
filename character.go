package main

import (
	"fmt"
	"math/rand"
	"time"

	vars "github.com/blaine-t-bush/wh3e-char-gen/vars"
)

func GenerateCharacter() Character {
	// Seed the randomizer.
	rand.Seed(time.Now().UnixNano())
	// Randomly select a class.
	class := GenerateClass()
	// Start at level 1.
	level := 1
	levelInfo := class.levels[level]
	// Randomly generate attributes.
	attributes := GenerateAttributeScores()
	// Generate hit points.
	hitPoints := GenerateHitPoints(levelInfo.hitDie, levelInfo.bonusHitPoints, level, class.name, attributes["con"].Score, 0)

	// Determine total group count.
	bonusGroupCount := 0
	for _, val := range attributes {
		if val.Score <= 5 {
			bonusGroupCount++
		}
	}
	remainingGroupCount := levelInfo.groupCount + bonusGroupCount

	// Determine species.
	species := GenerateSpecies()
	if species != "Human" {
		remainingGroupCount--
	}

	// Determine vocation.
	vocation := ""
	if remainingGroupCount > 0 {
		vocation = GenerateVocation()
		remainingGroupCount--
	}

	// Determine affiliations.
	var affiliations []string
	for {
		if remainingGroupCount <= 0 {
			break
		}

		affiliations = append(affiliations, GenerateAffiliation(species, affiliations))
		remainingGroupCount--
	}

	// Assign groups to attributes.
	var groups []Group

	if class.name != "Deft" {
		groups = append(groups, Group{
			name:      vocation,
			category:  "vocation",
			attribute: SelectGroupAttribute(groups, make([]string, 0)),
		})
	}

	if species != "Human" {
		firstSpeciesAttr := SelectGroupAttribute(groups, make([]string, 0))
		groups = append(groups, Group{
			name:      species,
			category:  "species",
			attribute: firstSpeciesAttr,
		})
		groups = append(groups, Group{
			name:      species,
			category:  "species",
			attribute: SelectGroupAttribute(groups, []string{firstSpeciesAttr}),
		})
	}

	for _, affiliation := range affiliations {
		groups = append(groups, Group{
			name:      affiliation,
			category:  "affiliation",
			attribute: SelectGroupAttribute(groups, make([]string, 0)),
		})
	}

	// Format groups for printing.
	groupsByAttribute := map[string][]string{
		"Strength":     {},
		"Dexterity":    {},
		"Constitution": {},
		"Intelligence": {},
		"Wisdom":       {},
		"Charisma":     {},
	}

	for _, group := range groups {
		groupsByAttribute[group.attribute] = append(groupsByAttribute[group.attribute], group.name)
	}

	attributeGroupStrings := map[string]string{
		"Strength":     "",
		"Dexterity":    "",
		"Constitution": "",
		"Intelligence": "",
		"Wisdom":       "",
		"Charisma":     "",
	}

	for attr, groupNames := range groupsByAttribute {
		if len(groupNames) > 0 {
			for index, groupName := range groupNames {
				if index == 0 {
					attributeGroupStrings[attr] += groupName
				} else {
					attributeGroupStrings[attr] += fmt.Sprintf(", %s", groupName)
				}
			}
		}
	}

	name := GenerateName(species)
	languages := GenerateLanguages(attributes["int"].Score)
	coins := GenerateCoins()

	formattedAttributes := map[string]Attribute{
		"str": {
			Score:  attributes["str"].Score,
			Groups: make([]string, 2),
		},
		"dex": {
			Score:  attributes["dex"].Score,
			Groups: make([]string, 2),
		},
		"con": {
			Score:  attributes["con"].Score,
			Groups: make([]string, 2),
		},
		"int": {
			Score:  attributes["int"].Score,
			Groups: make([]string, 2),
		},
		"wis": {
			Score:  attributes["wis"].Score,
			Groups: make([]string, 2),
		},
		"cha": {
			Score:  attributes["cha"].Score,
			Groups: make([]string, 2),
		},
	}

	character := Character{
		Name:        name,
		Level:       level,
		Class:       class.name,
		Species:     species,
		Vocation:    vocation,
		HitPoints:   hitPoints,
		AttackValue: levelInfo.attackValue,
		SavingThrow: levelInfo.savingThrow,
		Attributes:  formattedAttributes,
		Languages:   languages,
		Coins:       coins,
	}

	return character

	// fmt.Printf("%s, level %d %s %s %s\n", name, level, class.name, species, vocation)
	// fmt.Println("# Statistics")
	// fmt.Printf(" - HP  %2s\n", strconv.Itoa(hitPoints))
	// fmt.Printf(" - AV  %2s\n", strconv.Itoa(levelInfo.attackValue))
	// fmt.Printf(" - ST  %2s\n", strconv.Itoa(levelInfo.savingThrow))
	// fmt.Println("# Attributes")
	// fmt.Printf(" - STR %2s %s\n", strconv.Itoa(attributes["str"].score), attributeGroupStrings["Strength"])
	// fmt.Printf(" - DEX %2s %s\n", strconv.Itoa(attributes["dex"].score), attributeGroupStrings["Dexterity"])
	// fmt.Printf(" - CON %2s %s\n", strconv.Itoa(attributes["con"].score), attributeGroupStrings["Constitution"])
	// fmt.Printf(" - INT %2s %s\n", strconv.Itoa(attributes["int"].score), attributeGroupStrings["Intelligence"])
	// fmt.Printf(" - WIS %2s %s\n", strconv.Itoa(attributes["wis"].score), attributeGroupStrings["Wisdom"])
	// fmt.Printf(" - CHA %2s %s\n", strconv.Itoa(attributes["cha"].score), attributeGroupStrings["Charisma"])
	// fmt.Println("# Languages")
	// for _, language := range languages {
	// 	fmt.Printf(" - %s\n", language)
	// }
	// if class.name == "Deft" {
	// 	fmt.Println("# Attunements")
	// 	attunements := GenerateAttunements()
	// 	for index, attunement := range attunements {
	// 		if index == 0 {
	// 			fmt.Printf(" - %s*\n", attunement)
	// 		} else {
	// 			fmt.Printf(" - %s\n", attunement)
	// 		}
	// 	}
	// } else if class.name == "Strong" {
	// 	fmt.Println("# Abilities")
	// 	abilities := GenerateAbilities()
	// 	for _, ability := range abilities {
	// 		fmt.Printf(" - %s\n", ability)
	// 	}
	// } else if class.name == "Wise" {
	// 	fmt.Println("# Miracles")
	// 	miracles := GenerateMiracles(attributes["wis"].score)
	// 	for index, miracle := range miracles {
	// 		if index == 0 {
	// 			fmt.Printf(" - %s*\n", miracle)
	// 		} else {
	// 			fmt.Printf(" - %s\n", miracle)
	// 		}
	// 	}
	// }
	// fmt.Println("# Inventory")
	// fmt.Printf(" - %3s coins\n", strconv.Itoa(coins))
}

func GenerateName(species string) string {
	name := ""
	if species == "Dwarf" {
		name += GetRandomElement(vars.NamesDwarfAdjective) + " " + GetRandomElement(vars.NamesDwarfNoun)
	} else {
		prefixChance := 0.5
		rollPrefix := rand.Float64()

		// Randomly choose either prefix or suffix
		if rollPrefix <= prefixChance {
			name += GetRandomElement(vars.NamesHumanPrefix) + " " + GetRandomElement(vars.NamesHumanPrimary)
		} else {
			name += GetRandomElement(vars.NamesHumanPrimary) + " " + GetRandomElement(vars.NamesHumanSuffix)
		}
	}

	return name
}

func GenerateClass() Class {
	classNames := [3]string{
		"deft",
		"strong",
		"wise",
	}

	return Classes[classNames[rand.Intn(len(classNames))]]
}

func GenerateSpecies() string {
	// Normalize probabilities.
	var totalProbability float64 = 0
	for _, species := range Speciess {
		totalProbability += species.probability
	}

	// Randomly select a species.
	var cumulativeProbability float64 = 0
	var name string
	roll := rand.Float64()
	for _, species := range Speciess {
		cumulativeProbability += species.probability / totalProbability
		if roll <= cumulativeProbability {
			name = species.name
			break
		}
	}

	return name
}

// Randomly generate six attribute scores with 3d6.
func GenerateAttributeScores() map[string]Attribute {
	attributes := Attributes

	strength := attributes["str"]
	strength.Score = D(3, 6)
	attributes["str"] = strength

	dexterity := attributes["dex"]
	dexterity.Score = D(3, 6)
	attributes["dex"] = dexterity

	constitution := attributes["con"]
	constitution.Score = D(3, 6)
	attributes["con"] = constitution

	intelligence := attributes["int"]
	intelligence.Score = D(3, 6)
	attributes["int"] = intelligence

	wisdom := attributes["wis"]
	wisdom.Score = D(3, 6)
	attributes["wis"] = wisdom

	charisma := attributes["cha"]
	charisma.Score = D(3, 6)
	attributes["cha"] = charisma

	return attributes
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
	return GetRandomElement(vars.Vocations)
}

func GenerateAffiliation(species string, existingAffiliations []string) string {
	var affiliations []string
	if species == "Human" {
		affiliations = append(vars.Affiliations, vars.AffiliationsHuman...)
	} else if species == "Dwarf" {
		affiliations = append(vars.Affiliations, vars.AffiliationsDwarf...)
	}

	return GetUniqueElement(affiliations, existingAffiliations)
}

func GenerateAbility(existingAbilities []string) string {
	return GetUniqueElement(vars.Abilities, existingAbilities)
}

func GenerateAbilities() []string {
	var abilities []string
	abilityCount := 1

	for i := 0; i < abilityCount; i++ {
		abilities = append(abilities, GenerateAbility(abilities))
	}

	return abilities
}

func GenerateAttunement(existingAttunements []string) string {
	return GetUniqueElement(vars.Attunements, existingAttunements)
}

func GenerateAttunements() []string {
	var attunements []string
	attunementCount := 2

	for i := 0; i < attunementCount; i++ {
		attunements = append(attunements, GenerateAttunement(attunements))
	}

	return attunements
}

func GenerateMiracle(existingMiracles []string) string {
	return GetUniqueElement(vars.Miracles, existingMiracles)
}

func GenerateMiracles(wisdomScore int) []string {
	var miracles []string
	var miracleCount int

	if wisdomScore >= 16 {
		miracleCount = 4
	} else if wisdomScore >= 13 {
		miracleCount = 3
	} else {
		miracleCount = 2
	}

	for i := 0; i < miracleCount; i++ {
		miracles = append(miracles, GenerateMiracle(miracles))
	}

	return miracles
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
		languages = append(languages, GetUniqueElement(vars.Languages, languages))
	}

	return languages
}

func GenerateCoins() int {
	return D(3, 6) * 10
}
