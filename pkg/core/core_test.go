package core

import (
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		function func(int) int
		want     []int
	}{
		{"double values", []int{1, 2, 3}, func(x int) int { return x * 2 }, []int{2, 4, 6}},
		{"empty slice", []int{}, func(x int) int { return x * 2 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.slice, tt.function); !equal(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoldr(t *testing.T) {
	tests := []struct {
		name  string
		f     func(int, int) int
		acc   int
		slice []int
		want  int
	}{
		{"sum all elements", func(x, acc int) int { return x + acc }, 0, []int{1, 2, 3, 4}, 10},
		{"product of elements", func(x, acc int) int { return x * acc }, 1, []int{1, 2, 3, 4}, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foldr(tt.f, tt.acc, tt.slice); got != tt.want {
				t.Errorf("Foldr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoldl(t *testing.T) {
	tests := []struct {
		name  string
		f     func(int, int) int
		acc   int
		slice []int
		want  int
	}{
		{"sum all elements", func(acc, x int) int { return acc + x }, 0, []int{1, 2, 3, 4}, 10},
		{"product of elements", func(acc, x int) int { return acc * x }, 1, []int{1, 2, 3, 4}, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foldl(tt.f, tt.acc, tt.slice); got != tt.want {
				t.Errorf("Foldl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"min of 10 and 20", 10, 20, 10},
		{"min of -1 and 1", -1, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.a, tt.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"max of 10 and 20", 10, 20, 20},
		{"max of -1 and 1", -1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.a, tt.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to compare slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
