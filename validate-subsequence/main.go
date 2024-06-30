package main

import "fmt"

func validateSubSequence(arr, sub []int) bool {
	j := 0
	for _, val := range arr {
		if val == sub[j] {
			j++
		}
		if j == len(sub) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(validateSubSequence([]int{15, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}))
}
