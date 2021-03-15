package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	//进位
	carray := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carray
		if head == nil {
			head = &ListNode{Val: sum % 10}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum % 10}
			tail = tail.Next
		}
		carray = sum / 10
	}
	if carray > 0 {
		tail.Next = &ListNode{Val: carray}
	}
	return
}

func main() {
	l1 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 4, Next: l1}
	l3 := &ListNode{Val: 2, Next: l2}
	r1 := &ListNode{Val: 4}
	r2 := &ListNode{Val: 6, Next: r1}
	r3 := &ListNode{Val: 5, Next: r2}
	res := addTwoNumbers(l3, r3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
