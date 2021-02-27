package main

/* Race conditions explained

The main function launches 2 goroutines, numbered 1 and 2
Both goroutines count to a very high number and then set the shared 'winner' variable.
Whichever goroutine finishes first writes the value, but that is over-written by
whichever goroutine finishes last. (So, the 'winner' is the one that finishes last in this
case)

The code is in infinite loop, so you can see different results of multiple races.
i.e.,
Goroutine 1 set the winner variable
Goroutine 1 set the winner variable
Goroutine 2 set the winner variable
Goroutine 2 set the winner variable
Goroutine 1 set the winner variable
 */

import (
	"fmt"
	"sync"
)

func speak(num int, wg *sync.WaitGroup, winner *int) {
	for i := 0; i < 10000000; i++ {
	}
	*winner = num
	wg.Done()
}

func main() {
	for {
		var wg sync.WaitGroup
		winner := 0
		wg.Add(1)
		go speak(1, &wg, &winner)
		wg.Add(1)
		go speak(2, &wg, &winner)

		wg.Wait()
		fmt.Printf("Goroutine %d set the winner variable\n", winner)
	}
}
