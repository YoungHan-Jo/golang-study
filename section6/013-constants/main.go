package main

import "fmt"

const a = 42
const b = 42.78
const c = "James Bond" // constant's type is untyped constant

const (
	d = 42
	e = 42.78
	f = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d, e, f)

}
