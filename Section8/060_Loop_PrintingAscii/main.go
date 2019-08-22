package main

import "fmt"

func main() {
	fmt.Printf("Decimal\tUnicode\t\tHexadecimal\tText\n")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t\t%#U\t\t%#X\t\t%s\n", i, i, i, string(i))
	}
}
