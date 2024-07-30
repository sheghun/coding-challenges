package main

import "fmt"

func mergeOverlappingIntegers(arr *[][2]int) [][2]int {
	arr2 := [][2]int{}
	n := len(*arr)

	for i, j := 0, 1; j < n; i, j = i+1, j+1 {
		a := (*arr)[i]
		b := (*arr)[j]

		if a[1] >= b[0] || b[0] < a[0] {
			if a[1] < b[1] {
				a[1] = b[1]
				(*arr)[i] = a
			}
			if b[0] < a[0] {
				a[0] = b[0]
				(*arr)[i] = a
			}
			i--
			if n-j == 1 {
				arr2 = append(arr2, a)
				break
			}
		} else {
			arr2 = append(arr2, a)
			if n-j == 1 {
				arr2 = append(arr2, b)
				break
			}
			i = j - 1
		}
	}
	return arr2
}

func main() {
	// arr := [][2]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}}
	// arr := [][2]int{{1, 2}, {-20, 30}}
	arr := [][2]int{{43, 49}, {9, 12}, {12, 54}, {45, 90}, {91, 93}}
	fmt.Println(
		mergeOverlappingIntegers(
			&arr,
		),
	)
}
