package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestFindLastIndex(t *testing.T) {
	tests := []struct {
		name       string
		inputSlice []int
		callback   func(int) bool
		expected   int
	}{
		{
			"Empty slice",
			[]int{},
			func(i int) bool { return i == 10 },
			-1,
		},
		{
			"Value not found",
			[]int{1, 2, 3, 4, 5, 6, 7, 7},
			func(i int) bool { return i == 10 },
			-1,
		},
		{
			"Found the element",
			[]int{15, 6, 7, 72, 1, 2},
			func(i int) bool { return i > 6 },
			3,
		},
		{
			"Found the element at the end",
			[]int{15, 6, 7, 72, 1, 2},
			func(i int) bool { return i == 2 },
			5,
		},
		{
			"Found the element at the beginning",
			[]int{15, 6, 7, 72, 1, 2},
			func(i int) bool { return i == 15 },
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.FindLastIndex(tt.inputSlice, tt.callback)

			if !h.Equal(tt.expected, got) {
				t.Errorf("FindLastIndex - Want: %d but Got: %d", tt.expected, got)
			}
		})
	}
}
