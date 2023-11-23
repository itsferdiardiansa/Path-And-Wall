package main

import "fmt"

func main() {
	// var conferenceName = "Go Conference"
	// const conferenceTickets = 50
	// var remainingTickets = 10

	// fmt.Println("Welcom to booking application", conferenceName)
	// fmt.Printf("We have total of %v ticket(s) and %v ticket(s) still available", conferenceTickets, remainingTickets)
	// fmt.Println("Please pay attantion on every detail")

	// fmt.Printf("Hello %d\n", 23)
	// fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	// fmt.Println("Hello", 23)
	// fmt.Println(fmt.Sprint("Hello ", 23))

	// var x string = fmt.Sprintf("Hello %T %v \n", nil, 10)

	// fmt.Printf("Hallo %T %d \n", nil, 10)
	// fmt.Println(x)

	// printing.GetPrinting()
	// datatypes.PrintDataTypes()
	// pointer.RunPointer()
	// var num = 2
	// var status = true

	var (
		userName = "Ferdi Ardiansa"
		age      = 32
		city     = "Surabaya"
	)

	// const (
	// 	userActive   bool = false
	// 	lastActivity int  = 0
	// )

	// fmt.Printf("Type %v is %T\n", num, num)
	// fmt.Printf("Type %t is %T\n", status, status)

	for i := 0; i < len(userName); i++ {
		fmt.Println(userName[i], i)
	}

	fmt.Println(userName[0], age, city)
}
