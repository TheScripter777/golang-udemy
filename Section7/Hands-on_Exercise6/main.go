/**
Section 7 -- Hands-on exercise #4
Using iota, create 4 constants for the last 4 years. Print the constant values
*/
package main

import "fmt"

const (
	_ = iota
	a = 2019 - iota
	b = 2019 - iota
	c = 2019 - iota
	d = 2019 - iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
