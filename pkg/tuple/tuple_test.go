package tuple

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []string
		want   []Tuple[int, string]
	}{
		{
			name:   "equal length slices",
			slice1: []int{1, 2, 3},
			slice2: []string{"one", "two", "three"},
			want: []Tuple[int, string]{
				{First: 1, Second: "one"},
				{First: 2, Second: "two"},
				{First: 3, Second: "three"},
			},
		},
		{
			name:   "first slice longer",
			slice1: []int{1, 2, 3, 4},
			slice2: []string{"one", "two", "three"},
			want: []Tuple[int, string]{
				{First: 1, Second: "one"},
				{First: 2, Second: "two"},
				{First: 3, Second: "three"},
			},
		},
		{
			name:   "second slice longer",
			slice1: []int{1, 2},
			slice2: []string{"one", "two", "three", "four"},
			want: []Tuple[int, string]{
				{First: 1, Second: "one"},
				{First: 2, Second: "two"},
			},
		},
		{
			name:   "empty first slice",
			slice1: []int{},
			slice2: []string{"one", "two", "three"},
			want:   []Tuple[int, string]{},
		},
		{
			name:   "empty second slice",
			slice1: []int{1, 2, 3},
			slice2: []string{},
			want:   []Tuple[int, string]{},
		},
		{
			name:   "both slices empty",
			slice1: []int{},
			slice2: []string{},
			want:   []Tuple[int, string]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zip(tt.slice1, tt.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function min to calculate the minimum of two lengths
func min[T int](a, b T) T {
	if a < b {
		return a
	}
	return b
}
