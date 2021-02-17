package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortSlice(s []int) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func printSlice(s []int) {
	fmt.Printf("Sorted slice: len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var s []int
	s = make([]int, 0, 3)  // empty slice with capacity 3

	fmt.Println("Enter 'X' at any prompt to exit.")

	for {
		sortSlice(s)
		printSlice(s)
		var str string
		fmt.Println("Enter an intger:")
		str, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Printf(" >> Error %s: try again.", err)
			continue
		}

		// Ensure any embedded suffix / delimiter is removed as I'm using ReadString() function to capture.
		// In *nix, expect suffix '\n', in Windows expect suffix '\r\n'
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
		if str == "X" {
			fmt.Println("Exiting...")
			break
		}

		i, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf(" >> Error, not a integer: %s: try again.", err)
			continue
		}

		s = append(s, i)
	}
}