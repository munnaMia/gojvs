package helper

import "reflect"

// Checks two args are equal or not
func Equal[T any](x, y T) bool {
	return reflect.DeepEqual(x, y)
}

// Chechs the given number is a prime of or not
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}

	for factor := 3; factor*factor <= n; factor += 2 {
		if n%factor == 0 {
			return false
		}
	}

	return true
}
