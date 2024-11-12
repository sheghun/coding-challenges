package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m, t int
	ants := []int{}

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &t)

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		pos := x + (y * t)

		if pos > n {
			for pos > n {
				pos -= n
			}
		} else if pos <= 0 {
			for pos <= 0 {
				pos += n
			}
		}
		ants = append(ants, pos)
	}

	for i := 0; i < len(ants); i++ {
		for j := 0; j < len(ants)-1-i; j++ {
			if ants[j] > ants[j+1] {
				ants[j], ants[j+1] = ants[j+1], ants[j]
			}
		}
	}

	str := ""
	for _, val := range ants {
		str = fmt.Sprintf("%s %d", str, val)
	}

	fmt.Println(strings.Trim(str, " "))
}
