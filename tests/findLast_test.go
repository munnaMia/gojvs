package tests

import (
	"testing"

	"github.com/munnaMia/gojvs"
	h "github.com/munnaMia/gojvs/helper"
)

func TestFindLast(t *testing.T) {
	testsInt := []struct {
		name         string
		slice        []int
		expectedInt  int
		expectedBool bool
		callback     func(int) bool
	}{
		{
			"Found at start",
			[]int{1, 2, 3, 4, 5, 6, 7},
			7,
			true,
			func(a int) bool {
				return a == 7
			},
		},
		{
			"Not found",
			[]int{1, 2, 3, 4, 5, 6, 7},
			0,
			false,
			func(a int) bool {
				return a == 8
			},
		},
		{
			"Found at end",
			[]int{1, 2, 3, 4, 5, 6, 7},
			1,
			true,
			func(a int) bool {
				return a == 1
			},
		},
		{
			"Empty slice",
			[]int{},
			0,
			false,
			func(a int) bool {
				return a == 10
			},
		},
		{
			"Find prime number",
			[]int{12, 234, 5, 434, 42, 652, 210},
			5,
			true,
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
			gotInt, gotBool := gojvs.FindLast(tt.slice, tt.callback)

			if !h.Equal(gotInt, tt.expectedInt) {
				t.Errorf("Find() int - Want : %d & Got %d", tt.expectedInt, gotInt)
			}
			if !h.Equal(gotBool, tt.expectedBool) {
				t.Errorf("Find() bool - Want : %t & Got %t", tt.expectedBool, gotBool)
			}
		})
	}

	// string test
	testsString := []struct {
		name           string
		slice          []string
		expectedString string
		expectedBool   bool
		callback       func(string) bool
	}{
		{
			"Found at first index",
			[]string{"apple", "orenge", "pinapple"},
			"apple",
			true,
			func(str string) bool {
				return str == "apple"
			},
		},
		{
			"Not found",
			[]string{"apple", "orenge", "pinapple"},
			"",
			false,
			func(str string) bool {
				return str == "not found"
			},
		},
		{
			"Found at the last index",
			[]string{"apple", "orenge", "pinapple"},
			"pinapple",
			true,
			func(str string) bool {
				return str == "pinapple"
			},
		},
		{
			"Empty slice",
			[]string{},
			"",
			false,
			func(str string) bool {
				return str == "10"
			},
		},
	}

	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			gotInt, gotBool := gojvs.FindLast(tt.slice, tt.callback)

			if !h.Equal(gotInt, tt.expectedString) {
				t.Errorf("Find() string - Want : %s & Got %s", tt.expectedString, gotInt)
			}
			if !h.Equal(gotBool, tt.expectedBool) {
				t.Errorf("Find() bool - Want : %t & Got %t", tt.expectedBool, gotBool)
			}
		})
	}

}
