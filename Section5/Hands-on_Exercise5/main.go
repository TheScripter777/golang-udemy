/**
Section 5 -- Hands-on exercise #5
Building on the code from the previous example
	1. at the package level scope, using the var keyword, create a VARIABLE with the IDENTIFIER "y". The variable
	   should be of the UNDERLYING TYPE of your custom TYPE "x"
	2. in func main
		a. Now use the CONVERSION to convert the TYPE of the VALUE stored in "x" to the UNDERLYING TYPE
			i. Then use the short declaration operator to ASSIGN that value to "y"
*/
package main

import "fmt"

type exercise4 int
var x exercise4

// 1.
var y int

func main() {
	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Type of x: %T\n", x)
	x = 42
	fmt.Printf("New value of x: %v\n", x)

	// 2.
	y = int(x)

	fmt.Printf("Value of y: %v\n", x)
	fmt.Printf("Type of y: %T\n", x)
}
