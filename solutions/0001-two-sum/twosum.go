// 0001 >> two sum

package twosum

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, j := range nums {
		complement := target - j
		if res, ok := hashmap[complement]; ok {
			return []int{res, i}
		}
		hashmap[j] = i
	}

	return []int{}
}
