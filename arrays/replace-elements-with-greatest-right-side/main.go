package main

// leetcode problem 1299 replace elements with greatest element on the right side
// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/

import (
	"fmt"
	"math"
)

func replaceElements(arr []int) []int {
  var max int

	if len(arr) == 1 {
		arr[0] = -1
		return arr
	}

	for i, j := 0, 0; i < len(arr); i++ {
		if i == j {
			// get the biggest element to the right
      max = math.MinInt
			for k := j+1; k < len(arr); k++ {
				if arr[k] > max {
					j = k
					max = arr[k]
				}
			}
		}
		arr[i] = arr[j]
	}

  arr[len(arr) - 1] = -1

	return arr
}
func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
