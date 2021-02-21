package main

import (
	"fmt"
)

func getInts (ints *[]int) {
	var i int
	fmt.Println("Please enter up to 10 integers. A non-integer will stop entry.")
	for j := 1; j <= 10; j++ {
		fmt.Printf("Enter integer %d of 10: ", j)
		_, err := fmt.Scan(&i)
		if err != nil {
			fmt.Println("Non-integer input detected: Stopping user entry.")
			break
		}
		*ints = append(*ints, i)
	}
}

func Swap(ints []int, idx int) {
	temp := ints[idx]
	ints[idx] = ints[idx + 1]
	ints[idx + 1] = temp
}

func BubbleSort(ints []int) {
	numInts := len(ints)

	for endIdx := numInts - 1; endIdx > 0; endIdx-- {
		sorted := true
		for idx := 0; idx < endIdx; idx++ {
			e := ints[idx]
			if e > ints[idx+1] {
				Swap(ints, idx)
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}

func main() {
	var ints []int

	getInts(&ints)
	fmt.Printf("The original slice of integers: %v \n", ints)

	BubbleSort(ints)
	fmt.Printf("The sorted slice: %v \n", ints)
}