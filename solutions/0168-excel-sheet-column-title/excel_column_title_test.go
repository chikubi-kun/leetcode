package excelsheetcolumntitle

import "testing"

func TestConvertToTitle(t *testing.T) {
	type arg struct {
		input    int
		expected string
	}

	testCase := []arg{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
		{702, "ZZ"},
		{703, "AAA"},
	}

	for _, data := range testCase {
		if res := convertToTitle(data.input); res != data.expected {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
