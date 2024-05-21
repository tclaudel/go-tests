package mutant

import "testing"

func TestIsOfLegalAge(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected bool
	}{
		{
			name:     "Age 20",
			age:      20,
			expected: true,
		},
		{
			name:     "Age 10",
			age:      10,
			expected: false,
		},
		{
			name:     "Age 18",
			age:      18,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOfLegalAge(tt.age); got != tt.expected {
				t.Errorf("IsOfLegalAge() = %v, want %v", got, tt.expected)
			}
		})
	}
}
