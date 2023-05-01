package sliceutil

func Exists[A comparable](collection []A, item A) bool {
	for _, v := range collection {
		if v == item {
			return true
		}
	}

	return false
}
