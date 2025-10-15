package gojvs

import (
	"fmt"
)

// Map function takes an Slice and a callback function.
// It's modify slice element and return a new modified slice.
func Map[T any, R any](slice []T, fn func(T, int) R) []R {
	newSlice := make([]R, len(slice))

	for idx, value := range slice {
		newSlice[idx] = fn(value, idx)
	}

	return newSlice
}

// The filter() function creates a shallow copy of a portion
// of a given slice, filtered down to just the elements from
// the given slice that pass the test implemented by the
// provided function.
func Filter[T any](slice []T, fn func(T) bool) []T {
	filterSlice := make([]T, 0, cap(slice))
	for _, value := range slice {
		if ok := fn(value); ok {
			filterSlice = append(filterSlice, value)
		}
	}
	return filterSlice
}

// ForEach function takes an slice and a callback function with two arguments.
// Apply the callback function for each element in the slice.
// The callback function take a value and index like (T-> value, int-> index)
func ForEach[T any](slice []T, fn func(T, int)) {
	for idx, value := range slice {
		fn(value, idx)
	}
}

// Find Function takes an slice and a callback function and returns true & the first
// element in the provided array that satisfies the provided testing function else return
// false & a zerovalue.
func Find[T any](slice []T, fn func(T) bool) (T, bool) {
	var Zerovalue T

	for _, value := range slice {
		if ok := fn(value); ok {
			return value, true
		}
	}

	return Zerovalue, false

}

// The findIndex() function returns the index of the first element in an
// array that satisfies the provided testing function. If no elements
// satisfy the testing function, -1 is returned.
func FindIndex[T any](slice []T, fn func(T) bool) int {

	for idx, value := range slice {
		if ok := fn(value); ok {
			return idx
		}
	}

	return -1

}

// FindLast  function iterates the slice in reverse order. It takes an
// slice and a callback function and returns true & the first
// element in the provided array that satisfies the provided testing
// function else return false & a zerovalue.
func FindLast[T any](slice []T, fn func(T) bool) (T, bool) {
	var Zerovalue T

	for idx := len(slice) - 1; idx >= 0; idx-- {
		if ok := fn(slice[idx]); ok {
			return slice[idx], true
		}
	}

	return Zerovalue, false
}

// The FindLastIndex() function iterates the slice in reverse order and
// returns the index of the first element that satisfies the provided
// testing function. If no elements satisfy the testing function, -1 is
// returned.
func FindLastIndex[T any](slice []T, fn func(T) bool) int {

	for idx := len(slice) - 1; idx >= 0; idx-- {
		if ok := fn(slice[idx]); ok {
			return idx
		}
	}

	return -1

}

// The Concat function combines a collection of slices, flattening them into one
// unified slice.
func Concat[T any](slices ...[]T) []T {

	sliceContainer := make([]T, 0)
	for _, slice := range slices {
		sliceContainer = append(sliceContainer, slice...)
	}

	return sliceContainer
}

// The push() function  adds the specified elements to the end of an slice and returns
// the new length of the slice.
func Push[T any](slice *[]T, elems ...T) int {
	*slice = append(*slice, elems...)
	return len(*slice)
}

// The shift() Function removes the last element from an slice and returns that
// removed element.
func Pop[T any](slice *[]T) T {
	var lastElem T
	if len(*slice) == 0 {

		return lastElem
	}

	lastElem = (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]

	return lastElem

}

// The ToReverse() Function reverses an slice  and returns a new slice, the given
// slice element now becoming the last, and the last slice element becoming the
// first into the new slice.
func ToReverse[T any](slice []T) []T {
	length := len(slice)

	reverseSlice := make([]T, 0, length)

	for idx := length - 1; idx >= 0; idx-- {
		reverseSlice = append(reverseSlice, slice[idx])
	}

	return reverseSlice

}

// The shift() Function removes the first element from an slice and returns that
// removed element.
func Shift[T any](slice *[]T) T {
	var firstElem T
	if len(*slice) == 0 {

		return firstElem
	}

	firstElem = (*slice)[0]
	*slice = (*slice)[1:]

	return firstElem

}

// The some() function returns true if it finds one element in the slice that satisfies
// the provided testing function. Otherwise, it returns false.
func Some[T any](slice []T, fn func(T) bool) bool {
	for _, value := range slice {
		if ok := fn(value); ok {
			return true
		}
	}
	return false
}

// The copyWithin() function shallow copies part of this slice to another location in the
// same slice and returns this slice without modifying its length. to avoid panics we use
// explicit error handling. handle errors carefully.
func CopyWithIn[T any](slice *[]T, pos, start, end int) error {
	if pos < 0 {
		return fmt.Errorf("there is a value: %d, is less then: 0", pos)
	}
	if start < 0 {
		return fmt.Errorf("there is a value: %d, is less then: 0", start)
	}
	if end < 0 {
		return fmt.Errorf("there is a value: %d, is less then: 0", end)
	}

	// If the end position is greater then the length it will cause an error
	if end > len(*slice) {
		return fmt.Errorf("end position: %d, is greater then the actual length: %d of the slice", end, len(*slice))
	}

	if start > end {
		return fmt.Errorf("start position: %d is less then end position: %d", start, end)
	}

	// To avoid ref issue in slice create a new slice on a different location on the memory.
	tempSlice := make([]T, 0, end-start)

	for start < end {
		tempSlice = append(tempSlice, (*slice)[start])
		start++
	}

	for idx, value := range tempSlice {
		if idx+pos < len(*slice) {
			(*slice)[pos+idx] = value
		} else {
			return nil
		}
	}

	return nil
}

// The Fill() function changes all elements within a range of indices in an slice
// to a static value. And modified the actual slice.
func Fill[T any](slice *[]T, element T, start, end int) error {
	if start < 0 {
		return fmt.Errorf("start value: %d, is less then: 0", start)
	}
	if end < 0 {
		return fmt.Errorf("end value: %d, is less then: 0", end)
	}
	if start > end {
		return fmt.Errorf("start position: %d is less then end position: %d", start, end)
	}
	if end > len(*slice) {
		return fmt.Errorf("end position: %d, is greater then the actual length: %d of the slice", end, len(*slice))
	}

	for start < end {
		(*slice)[start] = element
		start++
	}

	return nil

}
