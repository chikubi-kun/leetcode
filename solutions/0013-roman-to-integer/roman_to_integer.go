package romantointeger

func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var out, prev int
	for _, char := range []byte(s) {
		val := dict[char]

		if prev < val {
			out -= (prev * 2)
		}

		out += val
		prev = val
	}

	return out
}
