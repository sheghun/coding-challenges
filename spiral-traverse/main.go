package main

import "fmt"

func spiralTraverse(arr [][]int) []int {
	m := len(arr)
	n := len(arr[0])

	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	u := 0
	path := []int{}

	seen := initSeen(m, n)

	x := 0
	y := 0

	for i := 0; i < m*n; i++ {
		path = append(path, arr[y][x])
		seen[y][x] = true

		tempx, tempy := x+dir[u][0], y+dir[u][1]

		if 0 <= tempy && tempy < m && 0 <= tempx && tempx < n && !seen[tempy][tempx] {
			x, y = tempx, tempy
		} else {
			u = (u + 1) % 4
			x, y = x+dir[u][0], y+dir[u][1]
		}
	}
	return path
}

func initSeen(m, n int) [][]bool {
	seen := [][]bool{}
	for i := 0; i < m; i++ {
		seen = append(seen, make([]bool, n))
	}
	return seen
}

func main() {
	fmt.Println(
		spiralTraverse([][]int{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		}),
	)
}
