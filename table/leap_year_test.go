package table_test

import (
	"github.com/stretchr/testify/assert"
	"go-tests/table"
	"testing"
	"time"
)

func TestLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{
			name: "2024 is a leap year",
			year: 2024,
			want: true,
		},
		{
			name: "2000 is a leap year",
			year: 2000,
			want: true,
		},
		{
			name: "1999 is not a leap year",
			year: 1999,
			want: false,
		},
		{
			name: "2004 is a leap year",
			year: 2004,
			want: true,
		},
		{
			name: "1900 is not a leap year",
			year: 1900,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			//
			time.Sleep(100 * time.Millisecond)

			//if got := table.LeapYear(tt.year); got != tt.want {
			//	t.Errorf("LeapYear() = %v, want %v", got, tt.want)
			//}

			assert.Equal(t, tt.want, table.LeapYear(tt.year))
		})
	}
}

// The package name is unittest_test. The package name is derived from the package being tested, unittest. Only Public functions can be tested.
// 1. The t.Run() function call in the test function is used to create a subtest.
// 2. The t.Parallel() function call in the test function tells the testing framework that the test function can be run in parallel with other test functions.
// 3. The time.Sleep(100 * time.Millisecond) function call in the test function simulates a long-running test.
// 4. t.Errorf() function call in the test function is used to report a test failure.
// 5. assert.Equal() equals the t.Errorf() function call in the test function. The assert.Equal() function is used to compare the actual and expected values.
