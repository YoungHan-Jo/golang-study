package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	bd := 1993
	for {
		if bd > 2023 {
			break
		}
		fmt.Println(bd)
		bd++
	}

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
