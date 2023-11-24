package main

import (
	"fmt"
	"strings"
)

func main() {
	// Conditional statement with "if"
	if name := "Ferdi"; strings.ToLower(name) == "ferdi" {
		fmt.Println("Match")
	} else {
		fmt.Println("Not found")
	}

	// Conditional statement with "switch"
	name := "Ferdi"

	// switch name := "Ferdi"; name == "ferdi" { // With short statement
	switch name {
	case "ferdi":
		fmt.Println("Match")
		break
	default:
		fmt.Println("Not found")
		break
	}

	// Without initial condition
	switch {
	case len(name) > 5:
		fmt.Println("Valid")
		break
	default:
		fmt.Println("Invalid")
		break
	}
}
