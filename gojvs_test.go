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
			// index i will never used on test cases. but you can bro if you want in any way.

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
