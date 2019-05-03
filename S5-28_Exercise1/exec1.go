package main

import "fmt"

func main() {
	exercise1()
}

/**
Section 5 -- 28. Hands-on exercise #1
	1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "x" and "y" and "z"
    	* 42
    	* "James Bond"
    	* true

	2. Now print the values stored in those variables using
		* A single print statement
		* Multiple print statements
*/
func exercise1() {
	// 1.
	x := 42
	y := "James Bond"
	z := true

	// 2.
	fmt.Println("Value of x: ", x)
	fmt.Println("Value of y: ", y)
	fmt.Println("Value of z: ", z)
	fmt.Printf("x = %v\ny = %v\nz = %v\n", x, y, z)
}
