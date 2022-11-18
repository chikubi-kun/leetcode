package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
	rows := make([]strings.Builder, numRows)
	pointer := 0
	upward := true

	for _, c := range s {
		rows[pointer].WriteRune(c)

		if pointer == 0 {
			upward = false
		}

		if pointer == numRows-1 {
			upward = true
		}

		if upward {
			pointer--
		} else {
			pointer++
		}
	}

	out := strings.Builder{}
	for _, row := range rows {
		out.WriteString(row.String())
	}

	return out.String()
}
