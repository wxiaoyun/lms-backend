package sliceutil

func Filter[S any](arr []S, predicate func(S) bool) []S {
	result := make([]S, 0, len(arr))
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
