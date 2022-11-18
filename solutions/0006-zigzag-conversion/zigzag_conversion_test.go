package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	type arg struct {
		input    string
		row      int
		expected string
	}

	testCase := []arg{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for _, data := range testCase {
		if res := convert(data.input, data.row); res != data.expected {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
