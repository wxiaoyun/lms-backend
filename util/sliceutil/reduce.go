package sliceutil

// Reduce reduces an array to a single value.
func Reduce[T any, R any](arr []T, reducer func(R, T) R, initialVal R) R {
	result := initialVal
	for _, v := range arr {
		result = reducer(result, v)
	}
	return result
}
