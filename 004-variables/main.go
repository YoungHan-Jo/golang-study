package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialization
var y = 43 // package level scope var declaration

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// assigns the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	x := 42 // short declaration operator in curly braces // use short declaration operator as you can
	fmt.Println(x)
	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
