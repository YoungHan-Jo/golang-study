package main

import "fmt"

func main() {
	x := 1

	for x <= 10 { // like while loop
		fmt.Println(x)
		x++
	}

	x = 1

	for { // infinite loop
		if x > 10 {
			break
		}
		if x%2 == 0 {
			x++
			continue
		}
		fmt.Println(x)
		x++
	}

	// init is executed once before condition is evaluated
	// for init; condition; post {}

}
