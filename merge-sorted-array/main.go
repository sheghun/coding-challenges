package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	left := m - 1
	right := n - 1

	for i := m + n - 1; i >= 0; i-- {
		if right < 0 {
			break
		}
		if left >= 0 && nums1[left] > nums2[right] {
			nums1[i] = nums1[left]
			left--
		} else {
			nums1[i] = nums2[right]
			right--
		}
	}
  fmt.Println(nums1)
}

func main() {
	merge([]int{2, 0}, 1, []int{1}, 1)
}
