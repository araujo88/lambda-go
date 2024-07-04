package utils

import (
	"reflect"
	"testing"
)

func TestHead(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"non-empty slice", []int{1, 2, 3}, 1},
		{"single element", []int{5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Head(tt.slice); got != tt.want {
				t.Errorf("Head() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTail(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"non-empty slice", []int{1, 2, 3}, []int{2, 3}},
		{"single element", []int{5}, []int{}},
		{"empty slice", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tail(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{"take more than exists", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"take none", []int{1, 2, 3}, 0, []int{}},
		{"take exact", []int{1, 2, 3}, 3, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Take(tt.slice, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{"drop none", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"drop all", []int{1, 2, 3}, 3, []int{}},
		{"drop some", []int{1, 2, 3}, 1, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Drop(tt.slice, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Drop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"reverse non-empty", []int{1, 2, 3}, []int{3, 2, 1}},
		{"reverse empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{"concat two slices", []int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3, 4, 5}},
		{"concat empty with non-empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.slice1, tt.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"unique elements", []int{1, 2, 2, 3, 4, 4, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"all unique", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"empty slice", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v, failed in test case: %s", got, tt.want, tt.name)
			}
		})
	}
}
