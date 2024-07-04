package utils

// head returns the first element from a slice
func Head[T any](slice []T) T {
	return slice[0]
}

// tail returns a slice without its first element
func Tail[T any](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	return slice[1:]
}

// returns the first n elements of the slice,
func Take[T any](slice []T, n int) []T {
	if len(slice) == 0 || n <= 0 {
		return []T{}
	}
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n]
}

func Drop[T any](slice []T, n int) []T {
	if len(slice) == 0 || n >= len(slice) {
		return []T{}
	}
	return slice[n:]
}

// reverses the order of a slice
func Reverse[T any](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	newSlice := make([]T, len(slice))
	for i, v := range slice {
		newSlice[len(slice)-1-i] = v
	}
	return newSlice
}

// concatenates two slices into one
func Concat[T any](slice1, slice2 []T) []T {
	return append(slice1, slice2...)
}

// returns a new slice containing unique elements from the original slice. This assumes elements are comparable
func Unique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	seen := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
