/**
Section 5 -- Hands-on exercise #2
	1. Use var to DECLARE three variables. The variables should have package level scope. Do not assign VALUES to the
		variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following type
		(meaning they can store values of that TYPE)
    		* Identifier "x" type int
    		* Identifier "y" type string
    		* Identifier "z" type bool

	2. Print out values for each identifier
	2.1 The compiler assigned values to the variables. What are these values called?
*/
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("Value of x: %v\nValue of y: %v\nValue of z: %v", x, y, z)

	// 2.1: Zero values
}
