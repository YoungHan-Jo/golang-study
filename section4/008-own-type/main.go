package main

import "fmt"

var a int

type customType int // underlying type of customType is int

var b customType

func main() {
	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b) // conversion, not casting in golang
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
