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

// The Concat function combines a collection of slices, flattening them into one unified slice.
func Concat[T any](slices ...[]T) []T {
	sliceContainer := make([]T, len(slices))
	for _, slice := range slices {
		sliceContainer = append(sliceContainer, slice...)
	}

	return sliceContainer
}
