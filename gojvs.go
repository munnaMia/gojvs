package gojvs

// Map function takes an Slice and a callback function.
// It's modify slice element and return a new modified slice.
func Map[T any, R any](slice []T, fn func(T) R) []R {
	newSlice := make([]R, len(slice))

	for idx, value := range slice {
		newSlice[idx] = fn(value)
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
