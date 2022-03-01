package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/linkovich2/dice"
)

type Special_Abilities struct {
	Name         string
	Desc         string
	Attack_Bonus int
	Damage_Dice  string
	Damage_Bonus int
}

type Monster struct {
	Name              string
	Challenge_Rating  string
	Dexterity         int
	Special_Abilities Special_Abilities
}

func doInitiative(monster1 Monster, monster2 Monster) {
	monsterPick1Initiative := rollSimpleInitiative(monster1)
	monsterPick2Initiative := rollSimpleInitiative(monster2)

	if monsterPick1Initiative >= monsterPick2Initiative {
		fmt.Printf("\n%s goes first", monster1.Name)
	} else {
		fmt.Printf("\n%s goes first", monster2.Name)
	}
}

func rollSimpleInitiative(monster Monster) int {
	dice.SeedRandom()
	var initiativeMod = (monster.Dexterity / 2) - 5
	var sbMonster1InitRoll strings.Builder
	var d1 dice.Dice = "1d20"

	sbMonster1InitRoll.WriteString("+")
	sbMonster1InitRoll.WriteString(strconv.Itoa(initiativeMod))
	initiative := d1.RollWithModifier(sbMonster1InitRoll.String())
	fmt.Printf(monster.Name + " rolled " + "(1d20" + sbMonster1InitRoll.String() + "): " + strconv.Itoa(initiative) + "\n")
	return initiative
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("monsters.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened monsters.json")
	//closing of our jsonFile so that we can parse it
	defer jsonFile.Close()

	//read our opened file as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Monsters array
	var monsters []Monster

	// we unmarshal our byteArray which contains our jsonFile's content into 'monsters' which we defined above
	json.Unmarshal(byteValue, &monsters)

	//printout each element for science!
	/*
		for k := range monsters {
			fmt.Printf("The monster '%s' has a CR of '%s' has an Initiative of '%d'\n", monsters[k].Name, monsters[k].Challenge_Rating, calcInitiative(monsters[k].Dexterity))
		}
	*/

	fmt.Printf("The  wizard summons 2 monsters out of %v: \n\n", len(monsters))
	randomMonster1 := rand.Intn(len(monsters))
	monsterPick1 := monsters[randomMonster1]
	randomMonster2 := rand.Intn(len(monsters))
	monsterPick2 := monsters[randomMonster2]

	//start fight sequence
	fmt.Printf("%s\nVS\n%s\n\n", monsterPick1.Name, monsterPick2.Name)

	//iniative phase
	fmt.Printf("---- Initiative Phase ----\n")
	doInitiative(monsterPick1, monsterPick2)

	//TODO finish round 1
}
