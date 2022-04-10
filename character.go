package main

import (
	"fmt"
	"strings"
)

type Class struct {
	id     int
	name   string
	baseXp int
	levels map[int]struct {
		xp int
		hd struct {
			base  int
			bonus int
		}
		attackValue int
		savingThrow int
		slotCount   int
		groupCount  int
		raiseCount  int
	}
}

type Group struct {
	name        string
	categoryId  int
	attributeId int
}

type GroupCategory struct {
	id int
	name string
}

type Attribute struct {
	id    int
	name  string
	score int
}

// GetAbbreviation returns the first three letters of an attribute in upper case.
func (attr Attribute) GetAbbreviation() string {
	return strings.ToUpper(attr.name[0:3])
}

type Character struct {
	name       string
	xp         int
	classId    int
	groups     []Group
	attributes [6]Attribute
}

var (
	classes = map[int]Class{
		1: {name: "Deft", baseXp: 1500},
		2: {name: "Strong", baseXp: 2000},
		3: {name: "Wise", baseXp: 2500},
	}

	groupCategories = map[int]GroupCategory{
		1:
	}

	attributes = map[int]Attribute{
		1: {name: "Strength"},
		2: {name: "Dexterity"},
		3: {name: "Constitution"},
		4: {name: "Intelligence"},
		5: {name: "Wisdom"},
		6: {name: "Charisma"},
	}
)

func main() {
	class := classes[1]
	fmt.Println("Name    :", class.name)
	fmt.Println("Base XP :", class.baseXp)

	for index, val := range attributes {
		fmt.Println("ID:          ", index)
		fmt.Println("Attribute:   ", val.name)
		fmt.Println("Abbreviation:", val.GetAbbreviation())
		fmt.Println("")
	}
}
