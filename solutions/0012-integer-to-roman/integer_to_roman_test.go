package integertoroman

import "testing"

func TestIntToRoman(t *testing.T) {
	type arg struct {
		input    int
		expected string
	}

	testCase := []arg{
		{input: 3, expected: "III"},
		{input: 58, expected: "LVIII"},
		{input: 1994, expected: "MCMXCIV"},
	}

	for _, data := range testCase {
		if res := intToRoman(data.input); res != data.expected {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
