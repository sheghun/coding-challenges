package main

import "fmt"

// leetcode problem 27
// https://leetcode.com/problems/remove-element/description/

func removeElement(nums []int, val int) int {
	k := 0
	for i, j := 0, len(nums)-1; i <= j; i++ {
		if nums[i] == val {
			// swap
			nums[i], nums[j] = nums[j], nums[i]
			// run this element again
			i--
			j--
		} else {
			k++
		}
	}

	return k
}

func main() {
	fmt.Println(removeElement([]int{1}, 1))
}
