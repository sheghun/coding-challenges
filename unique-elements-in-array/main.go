package main

import "fmt"

func UniQueElements(arr []int) []int {
	elements := make(map[int]int)
	arr2 := []int{}

	for _, val := range arr {
		if ok := elements[val]; ok > 0 {
			elements[val] += 1
		} else {
			elements[val] = 1
		}
	}

	for i, val := range elements {
		if val == 1 {
			arr2 = append(arr2, i)
		}
	}
	return arr2
}

func main() {
	fmt.Println(UniQueElements([]int{2, 3, 2, 4, 5, 6, 7, 8, 0, 0, 0, 9, 2}))
}
