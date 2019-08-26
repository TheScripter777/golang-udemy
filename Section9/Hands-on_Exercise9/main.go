/**
Section 9 -- Hands-on exercise #9
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string
with the IDENTIFIER "favSport"
*/
package main

import "fmt"

func main() {
	favSport := "lol"

	switch favSport {
	case "hockey":
		fmt.Println("Go Habs Go")

	case "fencing":
		fmt.Println("Duel anyone?")
	case "lol":
		fmt.Println("Welcome to summoner's rift!")
	}
}
