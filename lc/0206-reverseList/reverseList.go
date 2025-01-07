package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return h
}

func recursionReverseList(preNode *ListNode, curNode *ListNode) *ListNode {
	if curNode == nil {
		return preNode
	}

	nxtNode := curNode.Next
	curNode.Next = preNode

	return recursionReverseList(curNode, nxtNode)
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := recursionReverseList(head, head.Next)
	head.Next = nil

	return h
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preNode := head
	curNode := preNode.Next

	for curNode != nil {
		nxtNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nxtNode
	}

	head.Next = nil

	return preNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Printf("before reverse print list node: ")
	t := head
	for t != nil {
		fmt.Printf(" %d ", t.Val)
		t = t.Next
	}
	fmt.Printf("\r\n")

	head = reverseList(head)

	fmt.Printf("after reverse print list node: ")
	t = head
	for t != nil {
		fmt.Printf(" %d ", t.Val)
		t = t.Next
	}
	fmt.Printf("\r\n")

}
