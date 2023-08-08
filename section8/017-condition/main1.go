package main

import "fmt"

func main() {
	if 2 == 2 {
		fmt.Println("2 is equal to 2")
	}

	if x := 42; x != 2 {
		fmt.Println("initialization x is not equal to 2")
	}

}
