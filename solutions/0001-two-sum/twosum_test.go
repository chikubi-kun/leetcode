package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type arg struct {
		nums     []int
		target   int
		expected []int
	}

	testCase := []arg{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{nums: []int{1, 2, 3, 4}, target: 5, expected: []int{1, 2}},
	}

	for _, data := range testCase {
		if res := twoSum(data.nums, data.target); !reflect.DeepEqual(res, data.expected) {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
