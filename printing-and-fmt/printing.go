package printing

import (
	"fmt"
)

func GetPrinting() {
	/**
	 *
	 * Integer
	 *
	 * %b (base 2) binary
	 * %o (base 4) octal
	 * %d (base 10) decimal
	 * %x (base 16) hexadecimal
	 *
	 */

	fmt.Printf("Binary %b\n", 11)
	fmt.Printf("Oktal %o\n", 11)
	fmt.Printf("Decimal %d\n", 11)
	fmt.Printf("Hexadecimal %x\n", 11)
	fmt.Printf("Hexadecimal %X\n", 11)

	fmt.Printf("%9s\n", " ")
	fmt.Println("===========================")
	fmt.Printf("%9s\n", " ")

	/**
	 *
	 * Boolean
	 *
	 * %t(true or false)
	 *
	 */
	fmt.Printf("Boolean %t", true)
	fmt.Printf("Boolean %t", false)

	/**
	 *
	 * Floating
	 *
	 */
	fmt.Printf("Floating e %e\n", 2.35464756775466)
	fmt.Printf("Floating e %E\n", 2.35)
	fmt.Printf("Floating e %f\n", 2.35)
	fmt.Printf("Floating e %F\n", 2.354347364364364)
	fmt.Printf("Floating e %g\n", 2.354347364364364)
	fmt.Printf("Floating e %G\n", 2.354347364364364)

	fmt.Printf("%9s\n", " ")
	fmt.Println("===========================")
	fmt.Printf("%9s\n", " ")

	/**
	 *
	 * String
	 *
	 */
	fmt.Printf("This is about %s", "Golang\n")
	fmt.Printf("This is about %q \n", `
		And the value has been written by using a double-queted
	`)
}
