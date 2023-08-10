package sliceutil

func Map[S any, T any](arr []S, mapper func(S) T) []T {
	var result []T
	for _, v := range arr {
		result = append(result, mapper(v))
	}
	return result
}
