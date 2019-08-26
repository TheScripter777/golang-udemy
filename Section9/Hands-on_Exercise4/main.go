/**
Section 9 -- Hands-on exercise #4
Create a for loop using this syntax
for {}

Have it print out the year you have been alive
*/
package main

import "fmt"

func main() {
	bd := 1989
	for {
		if bd > 2019 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
