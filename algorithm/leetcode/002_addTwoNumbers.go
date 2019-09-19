package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry = 0
	var pre = &ListNode{
		-1,
		nil,
	}
	var cur = pre
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		a := 0
		b := 0
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val
		}

		if l2 == nil {
			b = 0
		} else {
			b = l2.Val
		}

		temp := a + b + carry
		carry = temp / 10
		cur.Next = &ListNode{temp % 10, nil}
		cur = cur.Next
		if l1 == nil {
			l1 = nil
		} else {
			l1 = l1.Next
		}
		if l2 == nil {
			l2 = nil
		} else {
			l2 = l2.Next
		}
	}
	return pre.Next
}
