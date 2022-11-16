package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// speed
func addtwonumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := l1
	carry := 0

	for {
		if l2 == nil {
			break
		}

		if l1.Next == nil {
			l1.Next = l2.Next
			l2.Next = nil
		}

		l1.Val += l2.Val + carry
		carry = l1.Val / 10
		l1.Val %= 10

		if l1.Next == nil {
			if carry > 0 {
				l1.Next = &ListNode{Val: 1}
			}
			return ans
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if carry == 0 {
		return ans
	}

	for {
		l1.Val += carry
		carry = l1.Val / 10
		l1.Val %= 10
		if l1.Next == nil {
			if carry > 0 {
				l1.Next = &ListNode{Val: 1}
			}
			return ans
		}
		l1 = l1.Next
	}
}

// recurrent
// func addtwonumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil && l2 == nil {
// 		return nil
// 	}

// 	if l1 == nil && l2 != nil {
// 		return l2
// 	}

// 	if l1 != nil && l2 == nil {
// 		return l1
// 	}

// 	next := addtwonumbers(l1.Next, l2.Next)
// 	sum := l1.Val + l2.Val

// 	if sum >= 10 {
// 		next = addtwonumbers(next, &ListNode{Val: sum / 10})
// 		sum %= 10
// 	}

// 	return &ListNode{Val: sum, Next: next}
// }

// func addtwonumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	t := 0
// 	head := new(*ListNode)
// 	current := head

// 	for l1 != nil || l2 != nil || t != 0 {
// 		if l1 != nil {
// 			t += l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			t += l2.Val
// 			l2 = l2.Next
// 		}
// 		if *current == nil {
// 			*current = new(ListNode)
// 		}
// 		val := t % 10
// 		t /= 10
// 		(**current).Val = val
// 		current = &((*current).Next)
// 	}
// 	return *head
// }
