package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Program should output "Found!" if user's string:
    1. starts with 'i'
    2. ends with 'n'
    3. contains 'a'
Otherwise, it should output "Not Found!"
 */

func checkStr(s string) bool {
	s = strings.ToLower(s)

	// Ensure any embedded suffix / delimiter is removed as I'm using ReadString() function to capture.
	// In *nix, expect suffix '\n', in Windows expect suffix '\r\n'
	s = strings.TrimSuffix(s, "\n")
	s = strings.TrimSuffix(s, "\r")

	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		return true
	}
	return false
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var s string

	fmt.Println("Enter a string:")
	s, err := stdin.ReadString('\n')

	if err != nil  {
		fmt.Printf("There was an error: %s \n", err)
		os.Exit(1)
	}

	if checkStr(s) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}