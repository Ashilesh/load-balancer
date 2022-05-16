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

func BinarySearch[T constraints.Ordered](arr []T, searchVal T) (int, bool) {
	start := 0
	end := len(arr) - 1

	for start < end {

		mid := (start + end) / 2

		if arr[mid] == searchVal {
			return mid, true
		}

		if arr[mid] > searchVal {
			end = mid - 1
		} else if arr[mid] < searchVal {
			start = mid + 1
		}
	}

	if arr[start] == searchVal {
		return start, true
	} else if searchVal < arr[start] {
		if start == 0 {
			return len(arr) - 1, false
		}
		return start - 1, false
	}

	return start, false
}
