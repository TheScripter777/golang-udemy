/**
Section 9 -- Hands-on exercise #3
Create a for loop using this syntax
for condition {}

Have it print out the year you have been alive
*/
package main

import "fmt"

func main() {
	bd := 1989
	for bd <= 2019 {
		fmt.Println(bd)
		bd++
	}
}