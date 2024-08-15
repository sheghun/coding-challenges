package main

import (
	"reflect"
	"testing"
)

func TestSplitNull(t *testing.T) {
	input := []any{1, 2, 3, 4, nil, 5, 6, 7, nil, 8, 9, 10}
	expected := [][]any{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}}
	result := SplitNull(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
