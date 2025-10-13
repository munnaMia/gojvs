package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestFilter(t *testing.T) {
	testsInt := []struct {
		name          string
		inputSlice    []int
		expectedSlice []int
		callback      func(int) bool
	}{
		{
			"All the even numbers",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{2, 4, 6},
			func(a int) bool {
				return a%2 == 0
			},
		},
		{
			"All positive numbers",
			[]int{1, 2, -3, 4, -5, 6, -7},
			[]int{1, 2, 4, 6},
			func(a int) bool {
				return a > 0
			},
		},
		{
			"Is prime number",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, -3, -2, -1, 8, 9, 10, 11, 12, 13},
			[]int{2, 3, 5, 7, 11, 13},
			h.IsPrime,
		},
		{
			"Empty slice",
			[]int{},
			[]int{},
			func(a int) bool {
				return true
			},
		},
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.Filter(tt.inputSlice, tt.callback)

			if !h.Equal(got, tt.expectedSlice) {
				t.Errorf("Find() int - Want : %v & Got %v", tt.expectedSlice, got)
			}

		})
	}

	// string test
	testsString := []struct {
		name          string
		inputSlice    []string
		expectedSlice []string
		callback      func(string) bool
	}{
		{
			"Minimum lenght is 6",
			[]string{"apple", "orenge", "pinapple"},
			[]string{"orenge", "pinapple"},
			func(str string) bool {
				return 6 <= len(str)
			},
		},
		{
			"Not match",
			[]string{"apple", "orenge", "pinapple"},
			[]string{},
			func(str string) bool {
				return str == "not found"
			},
		},
		{
			"Specific elements",
			[]string{"apple", "orenge", "pinapple"},
			[]string{"pinapple"},
			func(str string) bool {
				return str == "pinapple"
			},
		},
		
	}

	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.Filter(tt.inputSlice, tt.callback)

			if !h.Equal(got, tt.expectedSlice) {
				t.Errorf("Find() string - Want : %v & Got %v", tt.expectedSlice, got)
			}
			
		})
	}

}
