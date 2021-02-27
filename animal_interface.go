package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}
type Bird struct {}
type Snake struct {}

func (t Cow) Eat() {
		fmt.Println("grass")
	}

func (t Cow) Move() {
	fmt.Println("walk")
}

func (t Cow) Speak() {
	fmt.Println("moo")
}

func (t Bird) Eat() {
	fmt.Println("worms")
}

func (t Bird) Move() {
	fmt.Println("fly")
}

func (t Bird) Speak() {
	fmt.Println("chirp")
}

func (t Snake) Eat() {
	fmt.Println("mice")
}

func (t Snake) Move() {
	fmt.Println("slither")
}

func (t Snake) Speak() {
	fmt.Println("hsss")
}

func getUserCommand() (string, string, string) {
	stdin := bufio.NewReader(os.Stdin)
	var a, b, c string

	for {
		fmt.Print("> ")
		cmd, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error parsing your command, try again.")
			continue
		}
		fields := strings.Fields(cmd)
		a = fields[0]
		b = fields[1]
		c = fields[2]

		if len(fields) != 3 {
			fmt.Println("Error - did not find 3 tokens expected, try again.")
			continue
		}

		if a != "newanimal" && a != "query" {
			fmt.Println("Error - first string must be 'newanmimal' or 'query', try again.")
			continue
		}

		break
	}

	return a, b, c
}

func NewAnimal(kind string, name string) *Animal {
	var a Animal
	switch kind {
	case "cow":
		a = &Cow{}
	case "bird":
		a = &Bird{}
	case "snake":
		a = &Snake{}
	}
	_, _ = fmt.Printf("Created a new %s named %s!\n", kind, name)

	return &a
}

func main() {
	m := make(map[string]*Animal)

	for {
		cmd, name, s := getUserCommand()

		switch cmd {
		case "newanimal":
			kind := s
			m[name] = NewAnimal(kind, name)
		case "query":
			var a Animal
			a = *m[name]
			cmd := s
			switch cmd {
			case "eat":
				a.Eat()
			case "speak":
				a.Speak()
			case "move":
				a.Move()
			}

		}
	}


}
