/**
Section 7 -- 44. Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants
*/
package main

import "fmt"

const (
	x uint8 = 42
	y = 23
)

func main() {
	fmt.Printf("TYPED constant x (%T): %v\n", x, x)
	fmt.Printf("UNTYPED constant y (%T): %v\n", y, y)
}
