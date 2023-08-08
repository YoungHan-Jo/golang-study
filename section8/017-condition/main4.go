package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print1")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough // 次のcaseがfalseでも実行する
	case (7 == 3):
		fmt.Println("not true1")
		fallthrough
	case (3 == 31):
		fmt.Println("not true2")
		fallthrough
	case (31 == 31):
		fmt.Println("true 31")
	default:
		fmt.Println("this is default")
	}

	fmt.Println("========================================")

	switch n := "Moneypenny"; n {
	case "Moneypenny", "Bond", "Do No": // 複数のcaseを指定できる, どれか一つでもtrueなら実行される
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	default:
		fmt.Println("this is default")
	}
}
