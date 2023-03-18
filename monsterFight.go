package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// define the Monster struct
type Monster struct {
	Name        string `json:"name"`
	Size        string `json:"size"`
	Type        string `json:"type"`
	Subtype     string `json:"subtype"`
	Alignment   string `json:"alignment"`
	ArmorClass  int    `json:"armor_class"`
	HitPoints   int    `json:"hit_points"`
	HitDice     string `json:"hit_dice"`
	Speed       string `json:"speed"`
	Strength    int    `json:"strength"`
	Dexterity   int    `json:"dexterity"`
	Constitution int   `json:"constitution"`
	Intelligence int   `json:"intelligence"`
	Wisdom      int    `json:"wisdom"`
	Charisma    int    `json:"charisma"`
}

// define the RollDice function to simulate dice rolls
func RollDice(numDice int, numSides int) int {
	total := 0
	for i := 0; i < numDice; i++ {
		total += rand.Intn(numSides) + 1
	}
	return total
}

func main() {
	// initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// read the monsters JSON file
	monstersFile, err := ioutil.ReadFile("monsters.json")
	if err != nil {
		panic(err)
	}

	// parse the monsters JSON file into an array of Monster structs
	var monsters []Monster
	err = json.Unmarshal(monstersFile, &monsters)
	if err != nil {
		panic(err)
	}

	// randomly select two monsters
	monster1 := monsters[rand.Intn(len(monsters))]
	monster2 := monsters[rand.Intn(len(monsters))]

	// print the monsters that will fight
	fmt.Printf("%s vs. %s\n", monster1.Name, monster2.Name)

	// simulate the battle
	for {
		// monster1 attacks monster2
		damage := RollDice(1, 20)
		if damage >= monster2.ArmorClass {
			damage = RollDice(monster1.Dexterity/2, 6)
			monster2.HitPoints -= damage
			fmt.Printf("%s attacks %s for %d damage\n", monster1.Name, monster2.Name, damage)
		} else {
			fmt.Printf("%s misses %s\n", monster1.Name, monster2.Name)
		}

		// check if monster2 is defeated
		if monster2.HitPoints <= 0 {
			fmt.Printf("%s defeats %s\n", monster1.Name, monster2.Name)
			break
		}

		// monster2 attacks monster1
		damage = RollDice(1, 20)
		if damage >= monster1.ArmorClass {
			damage = RollDice(monster2.Dexterity/2, 6)
			monster1.HitPoints -= damage
			fmt.Printf("%s attacks %s for %d damage\n", monster2.Name, monster1.Name, damage)
		} else {
			fmt.Printf("%s misses %s\n", monster2.Name, monster1.Name)
		}

		// check if monster1 is defeated
		if monster1.HitPoints <= 0 {
			fmt.Printf("%s defeats %s\n", monster2.Name, monster1.Name)
			break
		}
	}
}
