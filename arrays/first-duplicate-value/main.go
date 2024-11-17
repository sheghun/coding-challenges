package main

import (
	"fmt"
	"math"
)

func firstDuplicate(arr []int) int {
	for _, num := range arr {
		i := int(math.Abs(float64(num))) - 1
		if arr[i] < 0 {
			return int(math.Abs(float64(num)))
		} else {
			arr[i] = arr[i] * -1
		}
	}
	return -1
}

func main() {
	fmt.Println(firstDuplicate([]int{2, 1, 5, 3, 3, 2, 4}))
}
