package main

import "fmt"

func checkIfExist(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}

	seen := make(map[int]int)

	for i, val := range arr {
		seen[val] = i
	}

	for i := 0; i < len(arr); i++ {
		val := arr[i] * 2
		if j, ok := seen[val]; ok && j != i {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}
