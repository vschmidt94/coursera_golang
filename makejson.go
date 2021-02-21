package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getInput(p string) string {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(p)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	return s
}


func main() {
	m := make(map[string]string)

	m["name"] = getInput("Name: ")
	m["address"] = getInput("Address: ")

	jsonP, _ := json.Marshal(m)

	fmt.Println("The JSON person object:")
	fmt.Println(string(jsonP))
}