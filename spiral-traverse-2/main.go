package main

import "fmt"

func spiralTraverse(arr [][]int) []int {
	var path []int

	n := len(arr[0]) - 1
	m := len(arr) - 1

	for top, right, bottom, left := 0, n, m, 0; left <= right && top <= bottom; top, right, bottom, left = top+1, right-1, bottom-1, left+1 {
		// left to right
		for i := left; i < right; i++ {
			path = append(path, arr[top][i])
		}

		// top to bottom
		for i := top; i < bottom; i++ {
			path = append(path, arr[i][right])
		}

		// right to left
		for i := right; i > left; i-- {
			path = append(path, arr[bottom][i])
		}

		// bottom to top
		for i := bottom; i > top; i-- {
			path = append(path, arr[i][left])
		}
	}

	return path
}

func main() {
	fmt.Println(
		spiralTraverse(
			[][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		),
	)

	fmt.Println(
		spiralTraverse([][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}),
	)
}
