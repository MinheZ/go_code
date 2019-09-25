package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var n = 10
	var ptr = 3
	head := CreateLoop(n)
	// ShowLoop(head, n)
	Josephus(head, ptr, n)
}

func Josephus(head *ListNode, ptr, n int) {
	if head == nil {
		fmt.Println("Empty loop!")
		return
	}
	ptr = n % ptr
	for head != head.Next {
		for i := 0; i < 1; i++ {
			head = head.Next
		}
		fmt.Printf("Removed No.%d\n", head.Next.Val)
		head.Next = head.Next.Next
	}
	fmt.Printf("Removed No.%d\n", head.Val)
}

func CreateLoop(n int) *ListNode {
	if n <= 0 {
		return nil
	}
	var dump = &ListNode{-1, nil}
	var cur = dump
	for i := 1; i <= n; i++ {
		cur.Next = &ListNode{i, nil}
		cur = cur.Next
	}
	cur.Next = dump.Next
	return dump.Next
}

func ShowLoop(head *ListNode, n int) {
	if head == nil {
		fmt.Println("Empty loop!")
	}
	var cur = head
	for i := 0; i < n; i++ {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
