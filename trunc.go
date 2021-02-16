package main

import (
	"bufio"
	"fmt"
	"os"
)

func clearBuffer(r *bufio.Reader) {
	_, _ = r.Discard(r.Buffered())
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var fp float32

	for {
		fmt.Println("Enter a floating point number:")
		_, err := fmt.Fscanln(stdin, &fp)
		if err == nil {
			t := int(fp)
			fmt.Printf("The truncated value is: %d \n", t)
			break
		} else {
			fmt.Printf("  >> There was an error: %s \n", err)
			fmt.Printf("Try again ... \n")
			clearBuffer(stdin)
		}
	}
}