package main

import "fmt"

func findSum(arr []int, targetSum int) []int {
	calculated := make(map[int]int)
	for _, b := range arr {
		a := targetSum - b
		// check if array element was already calculated
		if _, ok := calculated[b]; !ok {
			calculated[a] = b
			continue
		}

		// if it exists return the calculation and the element
		return []int{a, b}
	}
	return []int{0}
}

func main() {
	fmt.Println(findSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
}
