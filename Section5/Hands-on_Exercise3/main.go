/**
Section 5 -- 30. Hands-on exercise #3
Using the code from the previous exercise,
	1. At the package level scope, assign the following values to the three variables
		* for x assign 42
		* for y assign "James Bond"
		* for z assign true
	2. in func main
		* use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string
		  using the short declaration operator to a VARIABLE with the IDENTIFIER "s"

*/
package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)
}
