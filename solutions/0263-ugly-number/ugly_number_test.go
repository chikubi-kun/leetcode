package uglynumber

import "testing"

func TestIsUgly(t *testing.T) {
	type arg struct {
		input    int
		expected bool
	}

	testCase := []arg{
		{6, true},
		{1, true},
		{14, false},
	}

	for _, data := range testCase {
		if res := isUgly(data.input); res != data.expected {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}

}
