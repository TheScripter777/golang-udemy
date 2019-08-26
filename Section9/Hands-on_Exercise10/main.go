/**
Section 9 -- Hands-on exercise #10
Write down what these print:
true && true
true && false
true || true
true || false
!true
*/
package main

import "fmt"

func main() {
	fmt.Printf("true AND true ==> %t\n", true && true)
	fmt.Printf("true AND false ==> %t\n", true && false)
	fmt.Printf("true OR true ==> %t\n", true || true)
	fmt.Printf("true OR false ==> %t\n", true || false)
	fmt.Printf("NOT true ==> %t\n", !true)
}
