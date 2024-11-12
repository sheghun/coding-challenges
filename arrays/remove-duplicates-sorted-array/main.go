package main

// leetcode problem 27 remove duplicates from sorted array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	k := len(nums)
	if k == 1 {
		return k
	}

	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			k--
		} else {
			i++
			nums[i] = nums[j]
		}
	}

	return k
}

func main() {
}
