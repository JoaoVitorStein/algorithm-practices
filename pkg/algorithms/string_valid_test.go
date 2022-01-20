package algorithms

import "testing"

func TestIsStringValid(t *testing.T) {
	tests := []struct {
		name        string
		givenString string
		want        bool
	}{
		{
			name:        "valid string should return true",
			givenString: "{}",
			want:        true,
		},
		{
			name:        "valid string should return true",
			givenString: "[{()}]",
			want:        true,
		},
		{
			name:        "invalid string should return true",
			givenString: "[{()}]]",
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringValid(tt.givenString); got != tt.want {
				t.Errorf("IsStringValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
