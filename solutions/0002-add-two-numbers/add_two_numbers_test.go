package addtwonumbers

import (
	"reflect"
	"testing"
)

func createSingleLinkedList(arr []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, j := range arr {
		cur.Next = &ListNode{Val: j}
		cur = cur.Next
	}

	return head.Next
}

func TestAddTwoNumbers(t *testing.T) {
	testData := []struct {
		name     string
		arg1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name:     "one",
			arg1:     createSingleLinkedList([]int{2, 4, 3}),
			arg2:     createSingleLinkedList([]int{5, 6, 4}),
			expected: createSingleLinkedList([]int{7, 0, 8}),
		},
		{
			name:     "two",
			arg1:     createSingleLinkedList([]int{5}),
			arg2:     createSingleLinkedList([]int{5}),
			expected: createSingleLinkedList([]int{0, 1}),
		},
		{
			name:     "three",
			arg1:     createSingleLinkedList([]int{1}),
			arg2:     createSingleLinkedList([]int{9, 9, 9}),
			expected: createSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "four",
			arg1:     createSingleLinkedList([]int{9, 9, 9}),
			arg2:     createSingleLinkedList([]int{1}),
			expected: createSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "five",
			arg1:     createSingleLinkedList([]int{4, 3, 1}),
			arg2:     createSingleLinkedList([]int{1}),
			expected: createSingleLinkedList([]int{5, 3, 1}),
		},
		{
			name:     "six",
			arg1:     createSingleLinkedList([]int{1}),
			arg2:     createSingleLinkedList([]int{4, 3, 1}),
			expected: createSingleLinkedList([]int{5, 3, 1}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if result := addtwonumbers(data.arg1, data.arg2); !reflect.DeepEqual(result, data.expected) {
				t.Errorf("expected %v, got %v", data.expected, result)
			}
		})
	}
}
