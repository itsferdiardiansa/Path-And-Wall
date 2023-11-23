package main

import "fmt"

func main() {
	type IDNumber string

	var userID IDNumber = "1232323"
	var memberID int = 97

	fmt.Println(userID)
	fmt.Println(IDNumber(memberID))
}
