package gojvs

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
