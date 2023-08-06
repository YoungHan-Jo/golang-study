package main

import (
	"fmt"
	"runtime"
)

var x uint8 // 0 - 255
var y float64
var z int

func main() {
	x = 255
	y = 7.5
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
