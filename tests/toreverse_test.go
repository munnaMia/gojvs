package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestToReverse(t *testing.T) {
	tests := []struct {
		name          string
		inputSlice    []int
		expectedSlice []int
	}{
		{
			"Empty slice",
			[]int{},
			[]int{},
		},
		{
			"Slice with single element",
			[]int{22},
			[]int{22},
		},
		{
			"Slice with multiple element",
			[]int{22, 23, 24, 25, 26},
			[]int{26, 25, 24, 23, 22},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.ToReverse(tt.inputSlice)

			if !h.Equal(tt.expectedSlice, got) {
				t.Errorf("Reverse - Want: %v & Got: %v", tt.expectedSlice, got)
			}
		})
	}
}
