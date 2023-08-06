package main

import "fmt"

const (
	a = 2023 + iota
	b = 2023 + iota
	c = 2023 + iota
	d = 2023 + iota
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	x = x << 1
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
