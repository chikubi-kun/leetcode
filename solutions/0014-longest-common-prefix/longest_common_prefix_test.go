package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type arg struct {
		input    []string
		expected string
	}

	testCase := []arg{
		{input: []string{"flower", "flow", "flight"}, expected: "fl"},
		{input: []string{"dog", "racecar", "car"}, expected: ""},
	}

	for _, data := range testCase {
		if res := longestCommonPrefix(data.input); res != data.expected {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
