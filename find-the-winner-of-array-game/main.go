package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/find-the-winner-of-an-array-game/description/?source=submission-ac
func main() {
	fmt.Println(getWinner1([]int{3, 2, 1}, 10))
	fmt.Println(getWinner2([]int{3, 2, 1}, 10))
}

func getWinner1(arr []int, k int) int {
	wins := 0
	max := slices.Max(arr)

	if len(arr) < k {
		return max
	}

	for {
		i, j := arr[0], arr[1]

		if wins == k {
			return i
		}

		if i == max {
			return max
		}

		if i > j {
			wins++
			arr = append(arr[:1], arr[2:]...)
			arr = append(arr, j)
			continue
		}

		if i < j {
			wins = 1
			arr = arr[1:]
			arr = append(arr, i)
			continue
		}
	}
}

func getWinner2(arr []int, k int) int {
	wins := 0
	cur := arr[0]
	for i := 1; i < len(arr); i++ {
		if wins == k {
			return cur
		}

		if cur > arr[i] {
			wins++
			continue
		}
		cur = arr[i]
		wins = 1
	}

	return cur
}
