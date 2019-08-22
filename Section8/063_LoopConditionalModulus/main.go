package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is an even number (%%2==0)\n", i)
		}

		if i%3 == 0 {
			fmt.Printf("%d is a multiple of 3 (%%3==0)\n", i)
		}

		fmt.Println()
	}
}
