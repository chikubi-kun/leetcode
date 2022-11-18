package ranktransformofarray

import "sort"

// fastest > 77ms / 12.2 MB
func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	hashmap := make(map[int]int)
	sorted := make([]int, len(arr))

	copy(sorted, arr)
	sort.Ints(sorted)

	hashmap[sorted[0]] = 1
	for i, r := 1, 2; i < len(sorted); i++ {
		if sorted[i] != sorted[i-1] {
			hashmap[sorted[i]] = r
			r++
		}
	}

	for i := range arr {
		arr[i] = hashmap[arr[i]]
	}

	return arr
}

// lightest > 164ms / 10.8 MB
// func arrayRankTransform(arr []int) []int {
// 	hashmap := make(map[int]int)
// 	sorted := make([]int, len(arr))
// 	rank := 1

// 	copy(sorted, arr)
// 	sort.Ints(sorted)

// 	for _, val := range sorted {
// 		if _, ok := hashmap[val]; !ok {
// 			hashmap[val] = rank
// 			rank++
// 		}
// 	}

// 	for i := range arr {
// 		arr[i] = hashmap[arr[i]]
// 	}

// 	return arr
// }
