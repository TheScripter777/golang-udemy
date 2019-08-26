/**
Section 9 -- Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified
*/
package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Should NEVER print")
	case true:
		fmt.Println("True, then print it")
	}
}
