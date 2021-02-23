package main

import "fmt"

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

func getUserCommand() (string, string) {
	fmt.Print(">")
	return "cow", "speak"
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
