package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name          string
		inputSlices   [][]string
		expectedSlice []string
	}{
		{
			"Empty slices",
			[][]string{},
			[]string{},
		},
		{
			"Simple collection of slices",
			[][]string{
				{"a", "b", "c"},
				{"a", "b", "c"},
				{"a", "b", "c"},
			},
			[]string{"a", "b", "c", "a", "b", "c", "a", "b", "c"},
		},
		{
			"Mixed slices",
			[][]string{
				{"a", "b", "c"},
				{},
				{"ac", "bd", "ca"},
			},
			[]string{"a", "b", "c", "ac", "bd", "ca"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.Concat(tt.inputSlices...)

			if !h.Equal(tt.expectedSlice, got) {
				t.Errorf("Concat slices mismatch - Expected %v & Got %v", tt.expectedSlice, got)
			}
		})
	}
}
