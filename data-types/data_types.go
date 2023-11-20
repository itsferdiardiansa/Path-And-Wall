package datatypes

import "fmt"

func PrintDataTypes() {
	var userName string
	var age int
	var userStatus bool

	userName = "Ferdi Ardiansa"
	age = 32
	userStatus = false

	fmt.Printf("User %s has age %d and status %t \n", userName, age, userStatus)
}
