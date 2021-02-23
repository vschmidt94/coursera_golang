package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func getActions() []string {
	actions := make([]string, 0)
	actions = append(actions, "eat", "move", "speak")
	return actions
}

func getAnimals() []string {
	animals := make([]string, 0)
	animals = append(animals, "cow", "snamke", "bird")
	return animals
}

func stringInSlice(s string, sl []string) bool {
	for _, e := range sl {
		if e == s {
			return true
		}
	}
	return false
}

func getUserCommand() (string, string) {
	stdin := bufio.NewReader(os.Stdin)
	var action, animal string

	for {
		fmt.Print("> ")
		cmd, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error parsing your command, try again.")
			continue
		}
		fields := strings.Fields(cmd)
		animal = fields[0]
		action = fields[1]
		fmt.Println(animal, action)

		if len(fields) != 2 {
			fmt.Println("Error - did not find 2 tokens expected, try again.")
			printInstructions()
			continue
		}

		allowedAnimals := getAnimals()
		if !stringInSlice(animal, allowedAnimals) {
			fmt.Printf("Error - Animal %v not legal, must be one of %v \n", animal, allowedAnimals)
			continue
		}

		allowedActions := getActions()
		if !stringInSlice(action, allowedActions) {
			_, _ = fmt.Printf("Error - Action %v not legal, must be one of %v  \n", action, allowedActions)
			continue
		}

		break
	}

	return animal, action
}

func printInstructions() {
	fmt.Println("At the prompt, enter an animal followed by action.")
	fmt.Println("Animal choices are 'cow', 'bird', 'snake'")
	fmt.Println("Action choices are 'eat', 'move', 'speak'")
}


func main() {
	m := make(map[string]Animal)
	m["cow"] = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	m["bird"] = Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	m["snake"] = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	printInstructions()
	animal, action := getUserCommand()
	switch action {
	case "eat":
		m[animal].Eat()
	case "move":
		m[animal].Move()
	case "speak":
		m[animal].Speak()
	}

}
