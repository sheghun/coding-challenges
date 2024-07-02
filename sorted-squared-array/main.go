package main

import (
	"fmt"
	"math"
)

func sortedSquaredArray(arr []int) []int {
	sqrtArr := make([]int, len(arr))
	n := len(arr)
	right := n - 1
	left := 0

	for i := n - 1; i >= 0; i-- {
		leftVal := int(math.Abs(float64(arr[left])))
		rightVal := int(math.Abs(float64(arr[right])))
		if leftVal > rightVal {
			sqrtArr[i] = leftVal * leftVal
			left++
		} else {
			sqrtArr[i] = rightVal * rightVal
			right--
		}
	}

	return sqrtArr
}

func main() {
	fmt.Println(sortedSquaredArray([]int{-50, -13, -2, -1, 0, 0, 1, 1, 2, 3, 19, 20}))
	fmt.Println(sortedSquaredArray([]int{-7, -3, 1, 9, 22, 30}))
}
