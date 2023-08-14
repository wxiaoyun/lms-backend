package sliceutil

func Map[S any, T any](arr []S, mapper func(S) T) []T {
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		result = append(result, mapper(v))
	}
	return result
}
