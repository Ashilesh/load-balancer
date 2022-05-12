package utils

import (
	"golang.org/x/exp/constraints"
)

func Search[T constraints.Ordered](arr []T, searchVal T) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}
	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] > searchVal {
			break
		} else if arr[i] == searchVal {
			return i, true
		}
	}

	if i == 0 {
		return len(arr) - 1, false
	}
	return i - 1, false
}
