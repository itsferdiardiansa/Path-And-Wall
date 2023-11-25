package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) pushNotification(msg string) {
	fmt.Printf("%s %s\n", msg, customer.Name)
}

func main() {
	// var user = Customer{"Ferdi Ardiansa", "Surabaya", 32}
	user := Customer{
		Name:    "Ferdi Ardiansa",
		Address: "Surabaya",
		Age:     32,
	}

	// user.Name = "Ferdi Ardiansa"
	// user.Address = "Surabaya"
	// user.Age = 32

	fmt.Println(user)
	fmt.Printf("%T\n", user)

	user.pushNotification("A notification has been pushed to:")
}
