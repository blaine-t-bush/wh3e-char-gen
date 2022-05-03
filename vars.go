package main

type Class struct {
	name   string
	levels map[int]Level
}

type Level struct {
	xp             int
	hitDie         int
	bonusHitPoints int
	attackValue    int
	savingThrow    int
	slotCount      int
	groupCount     int
	raiseCount     int
}

type Group struct {
	name      string
	category  string
	attribute string
}

type GroupCategory struct {
	name           string
	attributeCount int
}

type Attribute struct {
	name         string
	abbreviation string
	score        int
}

type Character struct {
	name       string
	xp         int
	classId    int
	groups     []Group
	attributes [6]Attribute
}

var (
	attributes = map[string]Attribute{
		"str": {name: "Strength", abbreviation: "STR"},
		"dex": {name: "Dexterity", abbreviation: "DEX"},
		"con": {name: "Constitution", abbreviation: "CON"},
		"int": {name: "Intelligence", abbreviation: "INT"},
		"wis": {name: "Wisdom", abbreviation: "WIS"},
		"cha": {name: "Charisma", abbreviation: "CHA"},
	}

	groupCategories = map[string]GroupCategory{
		"vocation":    {name: "Vocation", attributeCount: 1},
		"affiliation": {name: "Affiliation", attributeCount: 1},
		"species":     {name: "Species", attributeCount: 2},
	}

	classes = map[string]Class{
		"deft": {
			name: "Deft",
			levels: map[int]Level{
				1: {
					xp:             0,
					hitDie:         1,
					bonusHitPoints: 0,
					attackValue:    10,
					savingThrow:    7,
					slotCount:      1,
					groupCount:     2,
					raiseCount:     0,
				},
				2: {
					xp:             1500,
					hitDie:         2,
					bonusHitPoints: 0,
					attackValue:    11,
					savingThrow:    8,
					slotCount:      1,
					groupCount:     2,
					raiseCount:     1,
				},
				3: {
					xp:             3000,
					hitDie:         2,
					bonusHitPoints: 1,
					attackValue:    11,
					savingThrow:    9,
					slotCount:      2,
					groupCount:     3,
					raiseCount:     1,
				},
				4: {
					xp:             6000,
					hitDie:         3,
					bonusHitPoints: 0,
					attackValue:    12,
					savingThrow:    10,
					slotCount:      2,
					groupCount:     3,
					raiseCount:     2,
				},
				5: {
					xp:             12000,
					hitDie:         3,
					bonusHitPoints: 1,
					attackValue:    12,
					savingThrow:    11,
					slotCount:      2,
					groupCount:     4,
					raiseCount:     2,
				},
				6: {
					xp:             24000,
					hitDie:         4,
					bonusHitPoints: 0,
					attackValue:    13,
					savingThrow:    12,
					slotCount:      2,
					groupCount:     4,
					raiseCount:     3,
				},
				7: {
					xp:             48000,
					hitDie:         4,
					bonusHitPoints: 1,
					attackValue:    13,
					savingThrow:    13,
					slotCount:      3,
					groupCount:     5,
					raiseCount:     3,
				},
				8: {
					xp:             96000,
					hitDie:         5,
					bonusHitPoints: 0,
					attackValue:    14,
					savingThrow:    14,
					slotCount:      3,
					groupCount:     5,
					raiseCount:     4,
				},
				9: {
					xp:             192000,
					hitDie:         5,
					bonusHitPoints: 1,
					attackValue:    14,
					savingThrow:    15,
					slotCount:      3,
					groupCount:     6,
					raiseCount:     4,
				},
				10: {
					xp:             384000,
					hitDie:         6,
					bonusHitPoints: 0,
					attackValue:    15,
					savingThrow:    16,
					slotCount:      4,
					groupCount:     6,
					raiseCount:     5,
				},
			},
		},
		"strong": {
			name: "Strong",
			levels: map[int]Level{
				1: {
					xp:             0,
					hitDie:         1,
					bonusHitPoints: 2,
					attackValue:    11,
					savingThrow:    5,
					slotCount:      1,
					groupCount:     2,
					raiseCount:     0,
				},
				2: {
					xp:             2000,
					hitDie:         2,
					bonusHitPoints: 0,
					attackValue:    11,
					savingThrow:    6,
					slotCount:      1,
					groupCount:     2,
					raiseCount:     1,
				},
				3: {
					xp:             4000,
					hitDie:         3,
					bonusHitPoints: 0,
					attackValue:    12,
					savingThrow:    7,
					slotCount:      1,
					groupCount:     2,
					raiseCount:     1,
				},
				4: {
					xp:             8000,
					hitDie:         4,
					bonusHitPoints: 0,
					attackValue:    13,
					savingThrow:    8,
					slotCount:      2,
					groupCount:     3,
					raiseCount:     2,
				},
				5: {
					xp:             16000,
					hitDie:         5,
					bonusHitPoints: 0,
					attackValue:    13,
					savingThrow:    9,
					slotCount:      2,
					groupCount:     3,
					raiseCount:     2,
				},
				6: {
					xp:             32000,
					hitDie:         6,
					bonusHitPoints: 0,
					attackValue:    14,
					savingThrow:    10,
					slotCount:      2,
					groupCount:     3,
					raiseCount:     3,
				},
				7: {
					xp:             64000,
					hitDie:         7,
					bonusHitPoints: 0,
					attackValue:    15,
					savingThrow:    11,
					slotCount:      3,
					groupCount:     4,
					raiseCount:     3,
				},
				8: {
					xp:             128000,
					hitDie:         8,
					bonusHitPoints: 0,
					attackValue:    15,
					savingThrow:    12,
					slotCount:      3,
					groupCount:     4,
					raiseCount:     4,
				},
				9: {
					xp:             256000,
					hitDie:         9,
					bonusHitPoints: 0,
					attackValue:    16,
					savingThrow:    13,
					slotCount:      3,
					groupCount:     4,
					raiseCount:     4,
				},
				10: {
					xp:             512000,
					hitDie:         10,
					bonusHitPoints: 0,
					attackValue:    17,
					savingThrow:    14,
					slotCount:      4,
					groupCount:     5,
					raiseCount:     5,
				},
			},
		},
		"wise": {
			name: "Wise",
			levels: map[int]Level{
				1: {
					xp:             0,
					hitDie:         1,
					bonusHitPoints: 1,
					attackValue:    10,
					savingThrow:    6,
					slotCount:      1,
					groupCount:     2,
					raiseCount:     0,
				},
				2: {
					xp:             2500,
					hitDie:         2,
					bonusHitPoints: 0,
					attackValue:    11,
					savingThrow:    7,
					slotCount:      1,
					groupCount:     2,
					raiseCount:     1,
				},
				3: {
					xp:             5000,
					hitDie:         2,
					bonusHitPoints: 1,
					attackValue:    11,
					savingThrow:    8,
					slotCount:      2,
					groupCount:     2,
					raiseCount:     1,
				},
				4: {
					xp:             10000,
					hitDie:         3,
					bonusHitPoints: 0,
					attackValue:    11,
					savingThrow:    9,
					slotCount:      2,
					groupCount:     3,
					raiseCount:     2,
				},
				5: {
					xp:             20000,
					hitDie:         4,
					bonusHitPoints: 0,
					attackValue:    12,
					savingThrow:    10,
					slotCount:      3,
					groupCount:     3,
					raiseCount:     2,
				},
				6: {
					xp:             40000,
					hitDie:         4,
					bonusHitPoints: 1,
					attackValue:    12,
					savingThrow:    11,
					slotCount:      3,
					groupCount:     3,
					raiseCount:     3,
				},
				7: {
					xp:             80000,
					hitDie:         5,
					bonusHitPoints: 0,
					attackValue:    12,
					savingThrow:    12,
					slotCount:      4,
					groupCount:     4,
					raiseCount:     3,
				},
				8: {
					xp:             160000,
					hitDie:         6,
					bonusHitPoints: 0,
					attackValue:    13,
					savingThrow:    13,
					slotCount:      4,
					groupCount:     4,
					raiseCount:     4,
				},
				9: {
					xp:             320000,
					hitDie:         6,
					bonusHitPoints: 1,
					attackValue:    13,
					savingThrow:    14,
					slotCount:      5,
					groupCount:     4,
					raiseCount:     4,
				},
				10: {
					xp:             640000,
					hitDie:         7,
					bonusHitPoints: 0,
					attackValue:    13,
					savingThrow:    15,
					slotCount:      5,
					groupCount:     5,
					raiseCount:     5,
				},
			},
		},
	}
)
