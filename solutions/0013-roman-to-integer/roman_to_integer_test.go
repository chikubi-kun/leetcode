package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	type arg struct {
		input    string
		expected int
	}

	testCase := []arg{
		{input: "III", expected: 3},
		{input: "LVIII", expected: 58},
		{input: "MCMXCIV", expected: 1994},
	}

	for _, data := range testCase {
		if res := romanToInt(data.input); res != data.expected {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
