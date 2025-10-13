package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestFindIndex(t *testing.T) {
	testsInt := []struct {
		name        string
		slice       []int
		expectedInt int
		callback    func(int) bool
	}{
		{
			"Found at beginning",
			[]int{1, 2, 3, 4, 5, 6, 7},
			0,
			func(a int) bool {
				return a == 1
			},
		},
		{
			"Not found",
			[]int{1, 2, 3, 4, 5, 6, 7},
			-1,
			func(a int) bool {
				return a == 8
			},
		},
		{
			"Found at end",
			[]int{1, 2, 3, 4, 5, 6, 7},
			6,
			func(a int) bool {
				return a == 7
			},
		},
		{
			"Empty slice",
			[]int{},
			-1,
			func(a int) bool {
				return a == 10
			},
		},
		{
			"Find prime number",
			[]int{12, 234, 5, 433, 42, 65, 210},
			2,
			func(a int) bool {
				if a < 2 || a%2 == 0 {
					return false
				}
				if a%2 == 0 {
					return a == 2
				}

				for factor := 3; factor*factor <= a; factor += 2 {
					if a%factor == 0 {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.FindIndex(tt.slice, tt.callback)

			if !h.Equal(got, tt.expectedInt) {
				t.Errorf("Find() int - Want : %d & Got %d", tt.expectedInt, got)
			}

		})
	}

	// string test
	testsString := []struct {
		name        string
		slice       []string
		expectedInt int
		callback    func(string) bool
	}{
		{
			"Found at beginning",
			[]string{"apple", "orenge", "pinapple"},
			0,
			func(str string) bool {
				return str == "apple"
			},
		},
		{
			"Not found",
			[]string{"apple", "orenge", "pinapple"},
			-1,
			func(str string) bool {
				return str == "not found"
			},
		},
		{
			"Found at end",
			[]string{"apple", "orenge", "pinapple"},
			2,
			func(str string) bool {
				return str == "pinapple"
			},
		},
		{
			"Empty slice",
			[]string{},
			-1,
			func(str string) bool {
				return str == "10"
			},
		},
	}

	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			got := gojvs.FindIndex(tt.slice, tt.callback)

			if !h.Equal(got, tt.expectedInt) {
				t.Errorf("Find() string - Want : %d & Got %d", tt.expectedInt, got)
			}
		})
	}

}
