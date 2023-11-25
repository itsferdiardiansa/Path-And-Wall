package main

import "fmt"

func notification() {
	fmt.Println("Notification has been invoked!")

	message := recover()

	if message != nil {
		fmt.Println("Something went wrong: ", message)
	}
}

func startApplication(e bool) {
	defer notification()

	if e {
		panic("Application failed to start:(")
	}

	fmt.Println("Apllication is starting")
}

func main() {
	startApplication(true)
}
