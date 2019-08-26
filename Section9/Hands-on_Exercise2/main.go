/**
Section 9 -- Hands-on exercise #2
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
U+0041 'A'
U+0041 'A'
U+0041 'A'

U+0042 'B'
U+0042 'B'
U+0042 'B'
...

*/
package main

import "fmt"

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		fmt.Println()
	}
	//otherSolution()
}

// i := 65 is equivalent to i := 'A' but IS NOT EQUIVALENT to i := "A"
func otherSolution() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		fmt.Println()
	}
}
