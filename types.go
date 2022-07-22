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
	Score        int
	Groups       []string
}

type Character struct {
	Name        string
	Level       int
	Class       string
	Species     string
	Vocation    string
	HitPoints   int
	SavingThrow int
	AttackValue int
	Attributes  map[string]Attribute
	Languages   []string
	Coins       int
}

type Species struct {
	name        string
	probability float64
}
