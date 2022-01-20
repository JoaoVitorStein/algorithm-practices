package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name       string
		givenArr   []int
		givenValue int
		want       int
	}{
		{
			name:       "when given a value that doesn't exist in array should return -1",
			givenArr:   []int{1, 2, 3, 4},
			givenValue: 5,
			want:       -1,
		},
		{
			name:       "when given a value that is in upper bound array should return its index",
			givenArr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			givenValue: 5,
			want:       4,
		},
		{
			name:       "when given a value that is in lower bound array should return its index",
			givenArr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			givenValue: 2,
			want:       1,
		},
		{
			name:       "when given a value that is in middle should return its index",
			givenArr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			givenValue: 4,
			want:       3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := BinarySearch(tt.givenArr, tt.givenValue); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
