package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	isCommon := func(common string, strs []string) bool {
		for _, word := range strs {
			if len(word) < len(common) || word[:len(common)] != common {
				return false
			}
		}
		return true
	}

	ln := len(strs[0])
	for {
		if ln == 0 {
			return ""
		}

		common := strs[0][:ln]
		if isCommon(common, strs) {
			return common
		}

		ln--
	}
}
