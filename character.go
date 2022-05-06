package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// Seed the randomizer.
	rand.Seed(time.Now().UnixNano())
	// Randomly select a class.
	class := GenerateClass()
	// Start at level 1.
	level := 1
	levelInfo := class.levels[level]
	// Randomly generate attributes.
	attributeScores := GenerateAttributeScores()
	// Generate hit points.
	hitPoints := GenerateHitPoints(levelInfo.hitDie, levelInfo.bonusHitPoints, level, 0)

	// Determine total group count.
	bonusGroupCount := 0
	for _, val := range attributeScores {
		if val <= 5 {
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

		affiliations = append(affiliations, GenerateAffiliation(affiliations))
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
			attributeGroupStrings[attr] = "("
			for index, groupName := range groupNames {
				if index == 0 {
					attributeGroupStrings[attr] += groupName
				} else {
					attributeGroupStrings[attr] += fmt.Sprintf(", %s", groupName)
				}
			}
			attributeGroupStrings[attr] += ")"
		}
	}

	name := GenerateName()

	fmt.Printf("%s, level %d %s %s %s\n", name, level, class.name, species, vocation)
	fmt.Println("# Statistics")
	fmt.Printf(" - HP  %2s\n", strconv.Itoa(hitPoints))
	fmt.Printf(" - AV  %2s\n", strconv.Itoa(levelInfo.attackValue))
	fmt.Printf(" - ST  %2s\n", strconv.Itoa(levelInfo.savingThrow))
	fmt.Println("# Attributes")
	fmt.Printf(" - STR %2s %s\n", strconv.Itoa(attributeScores["str"]), attributeGroupStrings["Strength"])
	fmt.Printf(" - DEX %2s %s\n", strconv.Itoa(attributeScores["dex"]), attributeGroupStrings["Dexterity"])
	fmt.Printf(" - CON %2s %s\n", strconv.Itoa(attributeScores["con"]), attributeGroupStrings["Constitution"])
	fmt.Printf(" - INT %2s %s\n", strconv.Itoa(attributeScores["int"]), attributeGroupStrings["Intelligence"])
	fmt.Printf(" - WIS %2s %s\n", strconv.Itoa(attributeScores["wis"]), attributeGroupStrings["Wisdom"])
	fmt.Printf(" - CHA %2s %s\n", strconv.Itoa(attributeScores["cha"]), attributeGroupStrings["Charisma"])
}
