package tests

import (
	"errors"
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestCopyWithIn(t *testing.T) {
	tests := []struct {
		name string

		inputSlices   []int
		inputPosition int
		startPosition int
		endPosition   int

		expectedSlice []int
		expectedErr   error
	}{
		{
			"Slices with negative position",
			[]int{2, 3, 4, 5, 6, 6},
			-1,
			1,
			1,
			[]int{2, 3, 4, 5, 6, 6},
			errors.New("there is a value: -1, is less then: 0"),
		},
		{
			"Slices with start position greater then end position",
			[]int{2, 3, 4, 5, 6, 6},
			0,
			10,
			1,
			[]int{2, 3, 4, 5, 6, 6},
			errors.New("start position: 10 is less then end position: 1"),
		},
		{
			"Slices with end position greated then length of slice",
			[]int{2, 3, 4, 5, 6, 6},
			0,
			1,
			10,
			[]int{2, 3, 4, 5, 6, 6},
			errors.New("end position: 10, is greater then the actual length: 6 of the slice"),
		},
		{
			"Simple collection of slices",
			[]int{2, 3, 4, 5, 6, 6},
			0,
			3,
			4,
			[]int{5, 3, 4, 5, 6, 6},
			nil,
		},
		{
			"Extended postions",
			[]int{11, 22, 33, 44, 55, 66},
			4,
			3,
			6,
			[]int{11, 22, 33, 44, 44, 55},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := gojvs.CopyWithIn(&tt.inputSlices, tt.inputPosition, tt.startPosition, tt.endPosition)

			if !h.Equal(tt.inputSlices, tt.expectedSlice) {
				t.Errorf("Concat slices - Want %v & Got %v", tt.expectedSlice, tt.inputSlices)
			}
			if !h.Equal(err, tt.expectedErr) {
				t.Errorf("Concat slices - Want %v & Got %v", tt.expectedErr, err)
			}
		})
	}
}
