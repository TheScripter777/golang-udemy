package main

import "fmt"

const (
	// Throw away the first iota because it's 0
	_ = iota

	// Shift 1 (1) ten times to the left == 1024 (10000000000)
	// Equivalent to: kb = 1024
	kb = 1 << (iota * 10) // iota = 1, shifts over 10
	mb = 1 << (iota * 10) // iota = 2, shifts over 20
	gb = 1 << (iota * 10) // iota = 3, shifts over 30
)

func main() {
	//x := 4
	//fmt.Printf("Var\t\tDec\t\tBin\n")
	//fmt.Printf("x\t\t%d\t\t%b\n", x, x)
	//
	//y := x << 1
	//fmt.Printf("y\t\t%d\t\t%b\n", y, y)
	//
	//z := x >> 1
	//fmt.Printf("z\t\t%d\t\t%b\n", z, z)
	//

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
