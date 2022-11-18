package ranktransformofarray

import (
	"reflect"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	type arg struct {
		input    []int
		expected []int
	}

	testCase := []arg{
		{[]int{40, 10, 20, 30}, []int{4, 1, 2, 3}},
		{[]int{100, 100, 100}, []int{1, 1, 1}},
		{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
	}

	for _, data := range testCase {
		if res := arrayRankTransform(data.input); !reflect.DeepEqual(res, data.expected) {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
