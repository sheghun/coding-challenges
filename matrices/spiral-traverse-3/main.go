package main

import "fmt"

func SpiralTraverse(arr [][]int) (arr2 []int) {
	if len(arr) == 0 {
		return []int{}
	}
	top := 0
	bottom := len(arr) - 1
	left := 0
	right := len(arr[0]) - 1

	if bottom == 0 {
		return arr[0]
	}

	for top <= bottom && left <= right {
		// left to right
		for i := left; i <= right; i++ {
			arr2 = append(arr2, arr[top][i])
		}
		top++
		// top to bottom
		for i := top; i <= bottom; i++ {
			arr2 = append(arr2, arr[i][right])
		}
		right--

		if top <= bottom {
			// right to left
			for i := right; i >= left; i-- {
				arr2 = append(arr2, arr[bottom][i])
			}
			bottom--

		}

		if left <= right {
			// bottom to top
			for i := bottom; i >= top; i-- {
				arr2 = append(arr2, arr[i][left])
			}
			left++
		}
	}
	return arr2
}

func main() {
	fmt.Println(SpiralTraverse(
		[][]int{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		},
	),
	)
}
