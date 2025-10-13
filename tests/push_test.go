package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name           string
		inputSlice     []int
		inputElems     []int
		expectedSlice  []int
		expectedLength int
	}{
		{
			"Empty slice",
			[]int{},
			[]int{},
			[]int{},
			0,
		},
		{
			"Push data into empty slice",
			[]int{},
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
			6,
		},
		{
			"Push data into a not empty slice",
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
			6,
		},
		{
			"Pushing no data",
			[]int{1, 2, 3},
			[]int{},
			[]int{1, 2, 3},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.Push(&tt.inputSlice, tt.inputElems...)

			if !h.Equal(tt.inputSlice, tt.expectedSlice) {
				t.Errorf("Push slice - Want: %v & Got: %v", tt.expectedSlice, tt.inputElems)
			}

			if !h.Equal(tt.expectedLength, got) {
				t.Errorf("Push length - Want: %d & Got: %d", tt.expectedLength, got)
			}
		})
	}
}
