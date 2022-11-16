package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	type arg struct {
		input    int
		expected bool
	}

	testCase := []arg{
		{input: 121, expected: true},
		{input: -121, expected: false},
		{input: 10, expected: false},
		{input: 1234567890987654321, expected: true},
	}

	for _, data := range testCase {
		if res := isPalindrome(data.input); res != data.expected {
			t.Errorf("expected %t, got %t", data.expected, res)
		}
	}
}
