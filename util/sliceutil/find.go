package sliceutil

func Find[S any](arr []S, predicate func(S) bool) *S {
	for _, v := range arr {
		if predicate(v) {
			return &v
		}
	}
	return nil
}
