package main

import "fmt"

func main() {
	x := 40

	if x == 40 {
		fmt.Println("x is equal to 40")
	} else if x == 41 {
		fmt.Println("x is equal to 41")
	} else {
		fmt.Println("x is not equal to 40 or 41")
	}

}
