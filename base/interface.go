package main

import "fmt"

type UserInterface interface {
	getName() string
}

type User struct {
	Firstname, Lastname, Email string
	isPremium                  bool
}

func (u User) getName() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

func printUsername(u UserInterface) {
	fmt.Println(u.getName())
}

// Any data types
func printRandomly() (any, any) {
	return 1, true
}

func main() {
	user := User{"Ferdi", "Ardiansa", "ff@mail.com", false}

	printUsername(user)

	fmt.Println(printRandomly())
}
