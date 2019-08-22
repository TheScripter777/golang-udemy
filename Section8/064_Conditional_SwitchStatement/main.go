package main

import "fmt"

func main() {
	var x int
	x = 3
	x++
	switch {
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("prints")
		fallthrough
	case x == 4:
		fmt.Println("also prints")
		fallthrough
	case x == 5:
		fmt.Println("false but prints anyway")
		fallthrough
	case x == 6, x % 2 ==0:
		fmt.Println("Multiple cases... true")
	}
}
