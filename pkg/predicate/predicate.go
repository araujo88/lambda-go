package predicate

func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// returns the first element in the slice that satisfies the specified predicate, or an indication that no such element exists.
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	var zero T // Zero value for type T
	return zero, false
}

// filters elements of a slice based on a predicate function, returning a new slice containing only elements that match the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var newSlice []T
	for _, v := range slice {
		if predicate(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// splits the slice into two slices based on a predicate: one for elements that satisfy the predicate and another for those that do not
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	var trueSlice []T
	var falseSlice []T
	for _, v := range slice {
		if predicate(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return trueSlice, falseSlice
}
