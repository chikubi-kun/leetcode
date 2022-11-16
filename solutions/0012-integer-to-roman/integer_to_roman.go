package integertoroman

import "strings"

type lmap struct {
	val int
	sym string
}

func intToRoman(num int) string {
	dict := []lmap{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	out := strings.Builder{}
	for _, e := range dict {
		for e.val <= num {
			num -= e.val
			out.WriteString(e.sym)
		}
	}

	return out.String()
}
