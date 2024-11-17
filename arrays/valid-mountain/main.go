package main

// leetcode problem 941 valid mountain array
// https://leetcode.com/problems/valid-mountain-array/description/

import "fmt"

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	for i, j, k := 0, 1, 2; k < n; i, j, k = i+1, j+1, k+1 {
		// check if it's valid peak
		if arr[i] < arr[j] && arr[k] < arr[j] {
			// go left
			for l := i; l >= 0; l-- {
				if l > -1 && arr[l] >= arr[l+1] {
					return false
				}
			}

			// go right
			for m := k; m < n; m++ {
				if m+1 < n && arr[m] <= arr[m+1] {
					return false
				}
			}

			return true
		}
	}

	return false
}

func main() {
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}
