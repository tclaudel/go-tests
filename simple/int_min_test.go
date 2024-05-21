package simple

import "testing"

func TestIntMin(t *testing.T) {
	if IntMin(2, -2) != -2 {
		t.Error("IntMin(2, -2) should be -2")
	}

	if IntMin(2, 0) != 0 {
		t.Error("IntMin(2, 0) should be 0")
	}

	if IntMin(-2, -2) != -2 {
		t.Error("IntMin(-2, -2) should be -2")
	}
	if IntMin(-2, -4) != -4 {
		t.Error("IntMin(-2, -4) should be -4")
	}

	if IntMin(5, 3) != 3 {
		t.Error("IntMin(5, 3) should be 3")
	}
}
