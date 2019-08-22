package main

import (
	"fmt"
	"time"
)

func foo() {

}

func main() {
	start := time.Now()
	foo()
	duration := time.Since(start)

	fmt.Println("Benchmark foo()...")
	fmt.Printf("Duration (ns) ==> %v\n", duration.Nanoseconds())
	fmt.Printf("Duration (sec) ==> %v\n", duration.Seconds())
}
