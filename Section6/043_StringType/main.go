package main

import "fmt"

func main() {
	s := "Hello, GoLang!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ---- %#X\n", s[i], s[i])
	}

	fmt.Println()
	fmt.Println()

	for i, v := range s {
		fmt.Printf("At index position %d, we have\n\thex %#x\n\tUTF-8 %#U\n", i, v, v)
	}
}
