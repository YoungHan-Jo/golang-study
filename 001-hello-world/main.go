package main

import "fmt"

func main() {
	fmt.Println("hello world")

	foo()

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
            continue
        }
		fmt.Print(i, " ")
	}

    bar();
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
    fmt.Println("and then we exited")
}
