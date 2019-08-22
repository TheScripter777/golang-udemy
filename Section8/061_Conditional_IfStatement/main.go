package main

import "fmt"

func main() {
	if x, y, z := 42, 42, "X and Y are equals"; x == y {
		fmt.Println(z)
	}
}
