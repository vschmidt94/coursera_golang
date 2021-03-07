package main

import (
	"fmt"
	"sort"
	"sync"
)

func getInts () []int {
	var i int
	var s []int
	fmt.Println("Please enter integers. A non-integer will stop entry.")
	for {
		fmt.Printf("Enter integer (X to stop): ")
		_, err := fmt.Scan(&i)
		if err != nil {
			if len(s) < 4 {
				fmt.Println("Need at least 4 integers before stopping.")
				continue
			}
			break
		}
		s = append(s, i)
	}
	return s
}

func sortInts (s *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(*s)
}

func combineTwoSlices(s1 []int, s2 []int, c chan []int) {
	var merged []int
	for {
		if len(s1) == 0 && len(s2) == 0 {
			break
		}
		if len(s1) == 0 && len(s2) > 0 {
			merged = append(merged, s2...)
			break
		}
		if len(s2) == 0 && len(s1) > 0 {
			merged = append(merged, s1...)
			break
		}
		if s1[0] < s2[0] {
			merged = append(merged, s1[0])
			s1 = s1[1:]
		} else {
			merged = append(merged, s2[0])
			s2 = s2[1:]
		}
	}
	c <- merged
}

func main() {
	var partitions [4][]int
	var wg sync.WaitGroup
	s := getInts()

	partitionLen := len(s) / 4
	fmt.Printf("Your original list of %d integers: %v\n", len(s), s)

	for i := 0; i < 4; i++ {
		if i != 3 {
			partitions[i] = s[i*partitionLen : (i+1)*partitionLen]
		} else {
			partitions[i] = s[i*partitionLen:]
		}

		fmt.Printf("Un-sorted partition %d: %v\n", i+1, partitions[i])
		wg.Add(1)
		go sortInts(&partitions[i], &wg)
	}

	wg.Wait()
	fmt.Println("All partitions have been individually sorted: ")
	for i := 0; i < 4; i++ {
		fmt.Println(partitions[i])
	}

	// combine slices in sorted order
	c := make(chan []int)
	go combineTwoSlices(partitions[0], partitions[1], c)
	go combineTwoSlices(partitions[2], partitions[3], c)
	s1, s2 := <-c, <-c
	go combineTwoSlices(s1, s2, c)
	sorted := <-c

	fmt.Println("The combined, sorted array:", sorted)





}