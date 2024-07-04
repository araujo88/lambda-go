package tuple

type Tuple[T any, U any] struct {
	First  T
	Second U
}

// zip combines elements of two slices into a slice of pairs (tuples), which is useful for combining related data.
func Zip[T any, U any](slice1 []T, slice2 []U) []Tuple[T, U] {
	length := min(len(slice1), len(slice2))
	result := make([]Tuple[T, U], length)
	for i := 0; i < length; i++ {
		result[i] = Tuple[T, U]{First: slice1[i], Second: slice2[i]}
	}
	return result
}
