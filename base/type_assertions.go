package main

import "fmt"

func random() any {
	return "OK!"
}

func main() {
	a := random()
	// aToString := a.(string)
	// fmt.Println(aToString)

	// b := random()
	// bToInt := b.(int) // panic: interface conversion: interface {} is string, not int
	// fmt.Println(bToInt)

	switch value := a.(type) {
	case string:
		fmt.Println("String", value)
		break
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
		break
	}
}
