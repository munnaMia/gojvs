package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestSome(t *testing.T) {
	testsInt := []struct {
		name       string
		inputSlice []int
		expected   bool
		callback   func(int) bool
	}{
		{
			"Is even numbers exist",
			[]int{1, 2, 3, 4, 5, 6, 7},
			true,
			func(a int) bool {
				return a%2 == 0
			},
		},
		{
			"Positive numbers",
			[]int{-1, -2, -3, 4, -5, 6, -7},
			true,
			func(a int) bool {
				return a > 0
			},
		},
		{
			"Eqult to zero",
			[]int{-1, -2, -3, 4, -5, 6, -7},
			false,
			func(a int) bool {
				return a == 0
			},
		},
		{
			"Is prime number",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, -3, -2, -1, 8, 9, 10, 11, 12, 13},
			true,
			h.IsPrime,
		},
		{
			"Empty slice",
			[]int{},
			false,
			func(a int) bool {
				return true
			},
		},
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.Some(tt.inputSlice, tt.callback)

			if !h.Equal(got, tt.expected) {
				t.Errorf("Find() int - Want : %t & Got %t", tt.expected, got)
			}

		})
	}
}
