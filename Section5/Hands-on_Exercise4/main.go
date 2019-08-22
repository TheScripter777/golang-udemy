/**
Section 5 Hands-on exercise #4
For this exercise
	1. Create your own type. Have the underlying type be an int.
	2. Create a VARIABLE of your new TYPE with the IDENTIFIDER "x" using the var keyword
	3. in func main
		a print out the value of the variable "x"
		b print out the type of the variable "x"
		c assign 42 to the VARIABLE "x" using the "=" OPERATOR
		d print out the value of the variable "x"


*/
package main

import "fmt"

// 1.
type exercise4 int

// 2.
var x exercise4

func main() {
	// 3a.
	fmt.Printf("Value of x: %v\n", x)

	// 3b.
	fmt.Printf("Type of x: %T\n", x)

	// 3c.
	x = 42

	// 3d.
	fmt.Printf("New value of x: %v\n", x)
}
