package main

import "fmt"

func longestPeak(arr []int) int {
	peak := 0

	if len(arr) < 3 {
		return 0
	}

	for i := 1; i < len(arr)-1; i++ {
		h := i - 1
		j := i + 1

		isPeak := arr[h] < arr[i] && arr[i] > arr[j]
		if !isPeak {
			continue
		}

		// expand left
		leftIdx := h
		for 0 < leftIdx && arr[leftIdx] > arr[leftIdx-1] {
			leftIdx--
		}

		rightIdx := j
		for len(arr)-1 > rightIdx && arr[rightIdx] > arr[rightIdx+1] {
			rightIdx++
		}

		currentPeakLength := rightIdx - leftIdx + 1

		if currentPeakLength > peak {
			peak = currentPeakLength
		}
	}

	return peak
}

func main() {
	fmt.Println(
		longestPeak([]int{
			1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3,
		}),
	)
}
