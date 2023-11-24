package main

import "fmt"

type Map map[string]interface{}

func main() {
	user := Map{
		"email":     "ferdi.ardiansa@yahoo.com",
		"isPremium": false,
	}

	fmt.Println(user)
	fmt.Println(user["emails"])

	book := make(map[string]string)

	book["name"] = "Golang"
	book["author"] = "JJ"
	book["date"] = "invalid"

	delete(book, "date")
	fmt.Println(book)
}
