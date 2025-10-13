package gojvs_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/munnaMia/gojvs"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name      string
		wantSlice any
		runTest   func() any
	}{
		{"Multiply by 2",
			[]int{2, 4, 6, 8, 100},
			func() any {
				slice := []int{1, 2, 3, 4, 50}
				callback := func(a int) int {
					return a * 2
				}
				return gojvs.Map(slice, callback)
			},
		},
		{"Empty slice", []int{}, func() any {
			slice := []int{}
			callback := func(a int) int {
				return a
			}
			return gojvs.Map(slice, callback)
		}},
		{"Converte Float32 to a Int number",
			[]int{1, 2, 3, 4, 5}, func() any {
				slice := []float32{1.33, 2.432, 3.4, 4.35, 5.0}
				callback := func(a float32) int {
					return int(a)
				}
				return gojvs.Map(slice, callback)
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.runTest()

			if !reflect.DeepEqual(got, tt.wantSlice) {
				t.Errorf("Map() got = %v, want %v", got, tt.wantSlice)
			}
		})
	}
}

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

			if !reflect.DeepEqual(actualSum, tt.expectedSum) {
				t.Errorf("ForEach Sum mismatch - Expected : %d & Got : %d", tt.expectedSum, actualSum)
			}

			if !reflect.DeepEqual(actualCount, tt.expectedCount) {
				t.Errorf("ForEach Count mismatch - Expected : %d & Got : %d", tt.expectedCount, actualCount)
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

			if !reflect.DeepEqual(formatedSlice, tt.expectedSlice) {
				t.Errorf("ForEach Slice mismatch - Expected : %v & Got %v", tt.expectedSlice, formatedSlice)
			}
		})
	}
}

func TestFind(t *testing.T) {
	testsInt := []struct {
		name         string
		slice        []int
		expectedInt  int
		expectedBool bool
		callback     func(int) bool
	}{
		{
			"Found at beginning",
			[]int{1, 2, 3, 4, 5, 6, 7},
			1,
			true,
			func(a int) bool {
				return a == 1
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
			7,
			true,
			func(a int) bool {
				return a == 7
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
			[]int{12, 234, 5, 433, 42, 65, 210},
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
			gotInt, gotBool := gojvs.Find(tt.slice, tt.callback)

			if !reflect.DeepEqual(gotInt, tt.expectedInt) {
				t.Errorf("Find() int mismatch - Expected : %d & Got %d", tt.expectedInt, gotInt)
			}
			if !reflect.DeepEqual(gotBool, tt.expectedBool) {
				t.Errorf("Find() bool mismatch - Expected : %t & Got %t", tt.expectedBool, gotBool)
			}
		})
	}

	// string test
	testsString := []struct {
		name         string
		slice        []string
		expectedInt  string
		expectedBool bool
		callback     func(string) bool
	}{
		{
			"Found at beginning",
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
			"Found at end",
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
			gotInt, gotBool := gojvs.Find(tt.slice, tt.callback)

			if !reflect.DeepEqual(gotInt, tt.expectedInt) {
				t.Errorf("Find() string mismatch - Expected : %s & Got %s", tt.expectedInt, gotInt)
			}
			if !reflect.DeepEqual(gotBool, tt.expectedBool) {
				t.Errorf("Find() bool mismatch - Expected : %t & Got %t", tt.expectedBool, gotBool)
			}
		})
	}

}
