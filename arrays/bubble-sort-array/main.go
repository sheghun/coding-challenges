package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	inputs := []int{}

	for i := 0; i < n; i++ {
		var d int

		if _, err := fmt.Scanf("%d", &d); err != nil {
			for err != nil {
				_, err = fmt.Scanf("%d", &d)
			}
			fmt.Println(d)
		}
		inputs = append(inputs, d)
	}

	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs)-1-i; j++ {
			if inputs[j] > inputs[j+1] {
				inputs[j], inputs[j+1] = inputs[j+1], inputs[j]
			}
		}
		fmt.Println(print(inputs))
	}
}

func print(inputs []int) string {
	str := ""
	for _, val := range inputs {
		str += " " + strconv.Itoa(val)
	}
	return strings.TrimSpace(str)
}
