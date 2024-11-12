package main

import (
	"reflect"
	"testing"
)

func TestSpiralTraverse(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{
				{1},
			},
			want: []int{1},
		},
		{
			matrix: [][]int{},
			want:   []int{},
		},
		{
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			want: []int{1, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SpiralTraverse(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
