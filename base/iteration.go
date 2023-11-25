package main

import (
	"fmt"
)

type Map map[string]interface{}

func longestSubsequence(str string, str2 string) int {
	graphs := [][]int{{0}}

	for i := 0; i < len(str2); i++ {
		graphs[0] = append(graphs[0], 0)
	}

	/**
	 *       s  a  y  a  b  a  c  a  k  i  o
	 *    0  0  0  0  0  0  0  0  0  0  0  0
	 *
	 *  b 0  0  0  0  0  1  1  1  1  1  1  1
	 *
	 *  a 0  0  1  1  1  1  2  2  2  2  2  2
	 *
	 *  c 0  0  1  1  1  1  1  3  3  3  3  3
	 *
	 *  a 0  0  1  1  2  2  2  3  4  4  4  4
	 *
	 *  i 0  0  1  1  2  2  2  3  4  4  5  5
	 *
	 *  o 0  0  1  1  2  2  2  3  4  4  5  6
	 *
	 *
	 */
	for i := 0; i < len(str); i++ {
		graphs = append(graphs, []int{0})

		for j := 0; j < len(str2); j++ {
			var temp []int

			if str[i] == str2[j] {
				temp = append(graphs[i+1], graphs[i][j]+1)
			} else {
				value := max(float64(graphs[i][j+1]), float64(graphs[i+1][j]))
				temp = append(graphs[i+1], int(value))

			}

			graphs[i+1] = temp

		}
	}

	fmt.Println(graphs)

	return -1
}

func main() {
	// name := "Ferdi"
	// index := 0

	// // Similar with while
	// for index < len(name) {
	// 	fmt.Printf("Index of %d: %s\n", index, string(name[index]))
	// 	index++
	// }

	// for i := 0; i < len(name); i++ {
	// 	fmt.Printf("Index of %d: %s\n", i, string(name[i]))
	// }

	// // For range
	// users := []string{"Ferdi", "Jake", "Matt"}
	// collections := []Map{
	// 	Map{
	// 		"username": "Ferdi Ardiansa",
	// 	},
	// 	Map{
	// 		"username": "Jake",
	// 	},
	// }

	// for i, value := range users {
	// 	fmt.Println(i, value)
	// }

	// fmt.Printf("%v %T\n", collections, collections)

	/**
	[
		[
			   b d e a v r a r r
			[0 0 0 0 0 0 0 0 0 0] 
			[0 1 1 1 1 1 1 1 1 1] 
			[0 1 1 1 2 2 2 2 2 2] 
			[0 1 1 1 2 2 3 3 3 3] 
			[0 1 1 1 2 2 3 4 4 4] 
			[0 1 1 1 2 2 3 4 5 5] 
			[0 1 1 1 2 2 3 4 5 5]
		]
	]
	*/
	fmt.Println(longestSubsequence("barara", "bdeavrarr"))
	fmt.Println(nil ?? 2)
}
