package main

import "fmt"

func getInts () []int {
	var i int
	var s []int
	fmt.Println("Please enter integers. A non-integer will stop entry.")
	for {
		fmt.Printf("Enter integer: ")
		_, err := fmt.Scan(&i)
		if err != nil {
			fmt.Println("Non-integer input detected: Stopping user entry.")
			break
		}
		s = append(s, i)
	}
}

func main() {
	s := getInts()

	var s1, s2, s3, s4 []int

}