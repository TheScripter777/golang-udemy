/**
Section 7 -- 45. Hands-on exercise #4
Write a program that
	* Assigns an int to a variable
	* Prints that int in decimal, binary and hex
	* Shifts the bits of that int over 1 position to the left, and assigns that to a variable
	* Prints that variable in decimal, binary and hex
*/
package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hexadecimal: %#x\n", x)

	fmt.Println()

	y := x << 1
	fmt.Printf("Decimal: %d\n", y)
	fmt.Printf("Binary: %b\n", y)
	fmt.Printf("Hexadecimal: %#x\n", y)

	fmt.Printf("\n\nBinary comparison\n")
	fmt.Printf("%d\t\t%b\n", x, x)
	fmt.Printf("%d\t\t%b\n", y, y)
}