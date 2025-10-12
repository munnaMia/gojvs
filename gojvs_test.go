package gojvs_test

import (
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
