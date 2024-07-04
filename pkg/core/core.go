package core

import "golang.org/x/exp/constraints"

// mapSlice applies a function to each element of a slice of type T and returns a new slice of the same type T.
func Map[T any](slice []T, function func(T) T) []T {
	newSlice := make([]T, len(slice))
	if len(slice) == 0 {
		return []T{}
	}
	for i, v := range slice {
		newSlice[i] = function(v)
	}
	return newSlice
}

// Foldr recursively folds a slice from the right using a function and an initial accumulator value.
func Foldr[T any, U any](f func(T, U) U, acc U, slice []T) U {
	if len(slice) == 0 {
		return acc
	}
	lastIndex := len(slice) - 1 // More efficient recursion
	return f(slice[lastIndex], Foldr(f, acc, slice[:lastIndex]))
}

// Foldr recursively folds a slice from the left using a function and an initial accumulator value.
func Foldl[T any, U any](f func(U, T) U, acc U, slice []T) U {
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}

// Min returns the smaller of two comparable values
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// max returns the greater of two comparable values
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
