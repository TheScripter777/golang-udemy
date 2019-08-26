package main

import "fmt"

func main() {
	var x int
	x = 3
	x++
	switch {
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("prints")
		fallthrough
	case x == 4:
		fmt.Println("also prints")
		fallthrough
	case x == 5:
		fmt.Println("false but prints anyway")
		fallthrough
	case x == 6, x % 2 ==0:
		fmt.Println("Multiple cases... true")
	}
}

// switch with no expression evaluates to true
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}

// no automatic fall through, but cases can be presented in comma-separeted lists
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}

	return false
}