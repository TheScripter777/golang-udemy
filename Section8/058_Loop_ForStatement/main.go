package main

import "fmt"

func main() {
	// Example of empty for
	// An empty 'for' is an eternal loop, breakable by 'break' command
	x := 1
	for {
		if x > 9 {
			break
		}

		fmt.Printf("x = %d\n", x)
		x++
	}
}
