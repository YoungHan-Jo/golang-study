package main

import "fmt"

const (
	x     = 42
	y int = 43
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
