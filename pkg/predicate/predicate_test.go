package predicate

import "testing"

func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{"true for positive match", []int{1, 2, 3}, func(x int) bool { return x == 2 }, true},
		{"false for no match", []int{1, 2, 3}, func(x int) bool { return x == 5 }, false},
		{"empty slice", []int{}, func(x int) bool { return x == 1 }, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.slice, tt.predicate); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{"true when all match", []int{2, 4, 6}, func(x int) bool { return x%2 == 0 }, true},
		{"false when one does not match", []int{2, 3, 6}, func(x int) bool { return x%2 == 0 }, false},
		{"empty slice returns true", []int{}, func(x int) bool { return x%2 == 0 }, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.slice, tt.predicate); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
		found     bool
	}{
		{"finds element", []int{1, 2, 3}, func(x int) bool { return x == 3 }, 3, true},
		{"does not find element", []int{1, 2, 3}, func(x int) bool { return x == 5 }, 0, false},
		{"empty slice", []int{}, func(x int) bool { return true }, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := Find(tt.slice, tt.predicate)
			if got != tt.want || found != tt.found {
				t.Errorf("Find() = %v, %v, want %v, %v", got, found, tt.want, tt.found)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{"filter out non-matching elements", []int{1, 2, 3}, func(x int) bool { return x > 1 }, []int{2, 3}},
		{"filter with no matches", []int{1, 2, 3}, func(x int) bool { return x == 5 }, []int{}},
		{"empty slice", []int{}, func(x int) bool { return true }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.slice, tt.predicate); !equal(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantTrue  []int
		wantFalse []int
	}{
		{"partition based on even", []int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }, []int{2, 4}, []int{1, 3}},
		{"empty slice", []int{}, func(x int) bool { return x%2 == 0 }, []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTrue, gotFalse := Partition(tt.slice, tt.predicate)
			if !equal(gotTrue, tt.wantTrue) || !equal(gotFalse, tt.wantFalse) {
				t.Errorf("Partition() gotTrue = %v, wantTrue %v; gotFalse = %v, wantFalse %v", gotTrue, tt.wantTrue, gotFalse, tt.wantFalse)
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
