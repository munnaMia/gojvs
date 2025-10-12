package gojvs

// Map function takes an Slice and a callback function.
func Map[T any, R any](slice []T, fn func(T) R) []R {
	newSlice := make([]R, len(slice))

	for idx, value := range slice {
		newSlice[idx] = fn(value)
	}

	return newSlice
}
