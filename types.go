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

// type Character struct {
// 	name        string
// 	class       Class
// 	groups      []Group
// 	attributes  map[string]int
// 	attackValue int
// 	savingThrow int
// 	armorClass  int
// 	moveSpeed   int
// }

type Species struct {
	name        string
	probability float64
}
