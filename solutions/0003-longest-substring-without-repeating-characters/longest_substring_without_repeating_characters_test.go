package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testData := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}

	expected := []int{3, 1, 3}

	for idx, data := range testData {
		if res := lengthOfLongestSubstring(data); res != expected[idx] {
			t.Errorf("expected %d, got %d", expected[idx], res)
		}
	}
}
