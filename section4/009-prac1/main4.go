package main

import "fmt"

type customType int

var x customType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 123
	fmt.Println(x)
}
