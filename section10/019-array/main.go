package main

import "fmt"

var x [5]int

func main() {
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[3])
	fmt.Println(x[3])
	fmt.Println(x[3])
	fmt.Println("upstream test")
}
