/**
Section 9 -- Hands-on exercise #7
Create a program that shows the "if statement" in action with "else if" and "else"
*/
package main

import "fmt"

func main() {
	x := "James Bond"

	if x == "Monneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Printf("BOND BOND")
	} else {
		fmt.Println("WHO?")
	}
}
