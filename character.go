package main

import (
	"fmt"
)

type class struct {
	id     int
	name   string
	baseXp int
}

type character struct {
	name    string
	classId int
}

var (
	classDeft = class{
		id:     1,
		name:   "Deft",
		baseXp: 1500,
	}
	
	classStrong = class{
		id:     2,
		name:   "Strong",
		baseXp: 2000,
	}
	
	classWise = class{
		id:     3,
		name:   "Wise",
		baseXp: 2500,
	}

	classes := [3]int{classDeft, classStrong, classWise}
)

func getLevel(classId string, xp int) int {

}

func newCharacter(name string) character {
	c := character{
		name:    name,
		classId: 1,
	}
	return c
}

func main() {
	c := newCharacter("Blaine")
	fmt.Println("Name ......", c.name)
	fmt.Println("Class ID ..", c.classId)
}
