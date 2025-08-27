package funcs

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	if l1 == nil {
		l1 = &ListNode{Val: 0}
	}
	if l2 == nil {
		l2 = &ListNode{Val: 0}
	}
	l3.Val = (l1.Val + l2.Val) % 10
	addNum := (l1.Val + l2.Val) / 10
	if l1.Next != nil || l2.Next != nil || addNum != 0 {
		l4 := AddTwoNumbers(l1.Next, l2.Next)
		l4.Val += addNum
		if l4.Val > 9 {
			l4.Val -= 10
			l5 := AddTwoNumbers(l4.Next, &ListNode{Val: 1})
			l4.Next = l5
		}
		l3.Next = l4
	}

	return l3
}

func OtherAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p0 := dummy
	m := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += m
		p0.Next = &ListNode{Val: sum % 10}
		m = sum / 10
		p0 = p0.Next
	}
	if m > 0 {
		p0.Next = &ListNode{Val: m}
	}
	return dummy.Next
}
