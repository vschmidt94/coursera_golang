package main

import "fmt"

func getValFor(s string) float64 {
	var f float64
	for {
		fmt.Printf("Enter %s: ", s)
		_, err := fmt.Scan(&f)
		if err == nil {
			break
		}
		fmt.Println(" > ERROR parsing your input, try again.")
	}

	return f
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	dispFn := func(t float64) float64 {
		s := 0.5 * a * t * t + v0 * t + s0
		return s
	}
	return dispFn
}

func main () {
	v0 := getValFor("initial velocity (v0)")
	s0 := getValFor("initial displacement (s0)")
	a := getValFor("acceleration (a)")
	t := getValFor("time in seconds (t)")

	fmt.Printf("")

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("After time %0.1f, the displacement is: %0.1f \n", t, fn(t))

}

