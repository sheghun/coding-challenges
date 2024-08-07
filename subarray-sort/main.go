package main

import "fmt"

func subArraySort(arr []int) []int {
	var a, b, s, e int

	s, e = -1, -1

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			a = arr[i]
			s = i
			break
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < arr[i-1] {
			b = arr[i]
			e = i
			break
		}
	}

	// expand left
	for i := s; i >= 0; i-- {
		if arr[i] <= a && s != i {
			s = i
			break
		}
	}

	// expand right
	for i := e + 1; i < len(arr); i++ {
		if arr[i] >= b && e != i {
			e = i
			break
		}
	}

	return []int{s, e}
}

func main() {
	fmt.Println(subArraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}))
	fmt.Println(subArraySort([]int{10, 12, 20, 30, 25, 40, 32, 31, 35, 50, 60}))
	fmt.Println(subArraySort([]int{-2, 0, 1, 15, 25, -1, 7, 30, 40, 50}))
}
