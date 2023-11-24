package main

import "fmt"

func reverseNumber(num int) int {
	result := 0

	for num > 0 {
		result = (result * 10) + (num % 10)

		// fmt.Println(result, num)
		num /= 10
	}

	return result
}

func main() {
	// nums := [...]int{1, 2, 3, 4, 5, 6}
	// nums2 := nums[:5]

	// nums[0] = 10

	// fmt.Println(nums, nums2, 123%10)

	// fmt.Println(reverseNumber(12345))

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	days2 := days[5:]

	days2[0] = "Sabtu"
	days2[1] = "Minggu"

	fmt.Println(days) // ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Sabtu", "Minggu"]

	days3 := append(days2, "Days off")
	fmt.Println(days2)      // ["Sabtu", "Minggu"]
	fmt.Println(days3)      // ["Sabtu", "Minggu", "Days off"]
	fmt.Println(cap(days3)) // 4

	// Make an initial slice with "make"
	newSlice := make([]string, 2, 5)

	newSlice[0] = "JKT"
	newSlice[1] = "SUB"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "SMG", "DNP", "MKS", "JYP", "SDJ", "PKL", "BND", "AMB", "PLB", "JGJ")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// Copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Initail declaration
	array := [...]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("Array %T, Slice %T\n", array, slice)
}
