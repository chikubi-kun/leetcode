package longestsubstringwithoutrepeatingcharacters

// import "strings"

func lengthOfLongestSubstring(s string) int {
	dict := [128]bool{}
	length, max := 0, 0

	for i, j := 0, 0; i < len(s); i++ {
		idx := s[i]
		for ; dict[idx]; j++ {
			length--
			dict[s[j]] = false
		}

		dict[idx] = true
		length++

		if length > max {
			max = length
		}
	}

	return max
}

// func lengthOfLongestSubstring(s string) int {
// 	var (
// 		start       = 0
// 		res         = 0
// 		lastOccured = make(map[string]int)
// 	)

// 	for i, j := range strings.Split(s, "") {
// 		if lastIndex, ok := lastOccured[j]; ok && lastIndex >= start {
// 			start = lastIndex + 1
// 		}

// 		if i-start+1 > res {
// 			res = i - start + 1
// 		}

// 		lastOccured[j] = i
// 	}

// 	return res
// }
