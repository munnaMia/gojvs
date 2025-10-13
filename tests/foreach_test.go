package tests

import (
	"fmt"
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestForEach(t *testing.T) {
	intTests := []struct {
		name          string
		slice         []int
		expectedSum   int
		expectedCount int
	}{
		{
			"Positive integer sum",
			[]int{1, 2, 3, 4, 5},
			15,
			5,
		},
		{
			"Empty slice sum",
			[]int{},
			0,
			0,
		},
		{
			"Positive and negative sum",
			[]int{10, -10, 30},
			30,
			3,
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			var actualSum, actualCount int

			// index i will never used on test cases. but you can use it bro if you want it in any way.
			callback := func(v, i int) {
				actualSum += v
				actualCount++
			}
			gojvs.ForEach(tt.slice, callback)

			if !h.Equal(actualSum, tt.expectedSum) {
				t.Errorf("ForEach Sum - Want : %d & Got : %d", tt.expectedSum, actualSum)
			}

			if !h.Equal(actualCount, tt.expectedCount) {
				t.Errorf("ForEach Count - Want : %d & Got : %d", tt.expectedCount, actualCount)
			}
		})
	}

	stringTests := []struct {
		name          string
		slice         []string
		expectedSlice []string
	}{
		{
			"Format empty slice",
			[]string{},
			[]string{},
		},
		{
			"Format single element slice",
			[]string{"Hello"},
			[]string{"0: Hello"},
		},
		{
			"Format string slice with many values",
			[]string{"Hello", "Hi", "Good", "To", "See", "U"},
			[]string{"0: Hello", "1: Hi", "2: Good", "3: To", "4: See", "5: U"},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			formatedSlice := make([]string, len(tt.slice))

			callback := func(str string, i int) {
				formatedSlice[i] = fmt.Sprintf("%d: %s", i, str)
			}

			gojvs.ForEach(tt.slice, callback)

			if !h.Equal(formatedSlice, tt.expectedSlice) {
				t.Errorf("ForEach Slice - Want : %v & Got %v", tt.expectedSlice, formatedSlice)
			}
		})
	}
}
