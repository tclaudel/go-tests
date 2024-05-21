package fuzz

import "testing"

func TestMustDivide(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Divide 10 by 2",
			a:        10,
			b:        2,
			expected: 5,
		},
		{
			name:     "Divide 12 by -3",
			a:        12,
			b:        -3,
			expected: -4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.a, tt.b); got != tt.expected {
				t.Errorf("Divide() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func FuzzMustDivide(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		Divide(a, b)
	})
}
