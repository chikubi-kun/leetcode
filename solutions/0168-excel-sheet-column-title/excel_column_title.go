package excelsheetcolumntitle

func convertToTitle(columnNumber int) string {
	var res []rune

	for columnNumber > 0 {
		mod := (columnNumber - 1) % 26
		res = append([]rune{rune('A' + mod)}, res...)
		columnNumber = (columnNumber - 1) / 26
	}

	return string(res)
}
