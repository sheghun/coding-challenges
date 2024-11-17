package main

import "fmt"

// https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/

func main() {
	fmt.Println(minSwaps([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}))
}

func minSwaps(grid [][]int) int {
	// calculate and storage row infos
	magic := make([]int, 0, len(grid))
loop1:
	for _, v := range grid {
		for i := len(v) - 1; i >= 0; i-- {
			if v[i] == 1 {
				magic = append(magic, i)
				continue loop1
			}
		}
		magic = append(magic, 0)
	}
	step := 0
loop2:
	for curr := 0; curr < len(grid); curr++ {
		for i, v := range magic {
			if v <= curr {
				copy(magic[1:i+1], magic[:i])
				magic = magic[1:]
				step += i
				continue loop2
			}
		}
		return -1
	}
	return step
}
