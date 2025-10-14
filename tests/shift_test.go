package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestShift(t *testing.T) {
	testsInt := []struct {
		name          string
		inputSlice    []int
		expectedSlice []int
		expectedInt   int
	}{
		{
			"Empty slice",
			[]int{},
			[]int{},
			0,
		},
		{
			"Single element",
			[]int{-7},
			[]int{},
			-7,
		},
		{
			"Two element",
			[]int{1, 2},
			[]int{2},
			1,
		},
		{
			"Multiple element in slice",
			[]int{2, 3, 4, 23, 53, 64, 15, 8, 69, 42, 5655, 5},
			[]int{3, 4, 23, 53, 64, 15, 8, 69, 42, 5655, 5},
			2,
		},
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.Shift(&tt.inputSlice)

			if !h.Equal(got, tt.expectedInt) {
				t.Errorf("Shift() int - Want : %d & Got %d", tt.expectedInt, got)
			}

			if !h.Equal(tt.inputSlice, tt.expectedSlice) {
				t.Errorf("Shift() slice - Want : %v & Got %v", tt.inputSlice, tt.expectedSlice)
			}

		})
	}
}
