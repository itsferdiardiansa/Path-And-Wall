package main

import (
	"fmt"
	"strings"
)

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)

	var studentName = "Ferdi"

	for i := 0; i < len(studentName); i++ {
		fmt.Println(strings.ToLower(string(studentName[i])))
	}
}
