package main

import "fmt"

func main() {
	var (
		names  [5]string
		cities = [3]string{
			"Jakarta",
			"Surabaya",
		}
		nums = [...]int{2, 4, 1, 3, 5, 12, 7}
	)

	names[0] = "Ferdi"
	names[1] = "Jack"
	names[2] = "Mike"
	names[3] = "Kevin"
	names[4] = "Jason"

	fmt.Println(names)
	fmt.Println(cities)

	nums[1] = 10
	fmt.Println(nums)

	// Bubble sort
	for i := 0; i < len(nums); i++ {
		for j := 0; j < (len(nums) - i - 1); j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j+1]

				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}

	fmt.Println(nums)
}
