package main

import "fmt"

func insertionSort(nums []int) (int, []int) {
	i := 1
	for i < len(nums) {
		j := i

		for (j-1) >= 0 && (nums[j-1] > nums[j]) {
			temp := nums[j-1]

			nums[j-1] = nums[j]
			nums[j] = temp

			j--
		}

		i++
	}

	return len(nums), nums
}

// Declare variables as return of function
func getName(fs string, ls string) (firstName, lastName string) {
	firstName = "This is " + fs
	lastName = "This is " + ls
	return
}

// Variadic function
func sumAll(nums ...int) (result int) {
	for _, val := range nums {
		result += val
	}

	return
}

// Recursive
func recursive(n int) {
	fmt.Println(n)

	if n <= 0 {
		return
	}

	recursive(n - 1)
}

type Callback func(int, string)

func checkout(amount int, successCalback Callback) {
	successCalback(amount, "Checkout completed!")
}

// Closure
func increment() func() int {
	counter := 0

	return func() int {
		counter++

		return counter
	}
}

func main() {
	p := insertionSort
	_, sorted := insertionSort([]int{5, 4, 1, 7, 3, 2, 10})

	fmt.Println(sorted)
	fmt.Println(getName("Ferdi", "Ardiansa"))
	fmt.Printf("%T\n", p)

	fmt.Println(sumAll(1, 2, 3))

	// recursive(5)

	checkout(2000, func(t int, s string) {
		fmt.Printf("Your products cost is: %d \n", t)
		fmt.Printf("%s \n", s)
	})

	c := increment()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}
