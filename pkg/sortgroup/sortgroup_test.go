package sortgroup

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		keyFunc func(int) string
		want    map[string][]int
	}{
		{
			name:  "group by parity",
			slice: []int{1, 2, 3, 4, 5},
			keyFunc: func(x int) string {
				if x%2 == 0 {
					return "even"
				}
				return "odd"
			},
			want: map[string][]int{
				"odd":  {1, 3, 5},
				"even": {2, 4},
			},
		},
		{
			name:    "empty slice",
			slice:   []int{},
			keyFunc: func(x int) string { return "odd" },
			want:    map[string][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupBy(tt.slice, tt.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"sort positive numbers", []int{5, 3, 4, 1, 2}, []int{1, 2, 3, 4, 5}},
		{"sort including negatives", []int{-1, -3, 2, 1, 0}, []int{-3, -1, 0, 1, 2}},
		{"empty slice", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortSlice(tt.slice) // Modify the slice in-place
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("sortSlice() resulted in %v, want %v", tt.slice, tt.want)
			}
		})
	}
}
