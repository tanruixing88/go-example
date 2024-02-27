package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	var preStartNode *ListNode

	startNode := head
	endNode := startNode
	curCount := 1

	for endNode != nil {
		if curCount < k {
			curCount++
			endNode = endNode.Next
		} else {
			lNode := startNode
			rNode := endNode
			endNode = endNode.Next
			startNode = endNode
			curCount = 1

			cNode := lNode
			nNode := cNode.Next
			for nNode != rNode {
				nnNode := nNode.Next
				nNode.Next = cNode
				cNode = nNode
				nNode = nnNode
			}

			lNode.Next = rNode.Next
			rNode.Next = cNode

			if preStartNode == nil {
				head = rNode
			} else {
				preStartNode.Next = rNode
			}

			preStartNode = lNode
		}
	}

	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	fmt.Printf("before reverse: ")
	t := head
	for t != nil {
		fmt.Printf(" %d", t.Val)
		t = t.Next
	}
	fmt.Printf("\r\n")
	head = reverseKGroup(head, 3)
	fmt.Printf("after reverse: ")
	t = head
	for t != nil {
		fmt.Printf(" %d", t.Val)
		t = t.Next
	}
	fmt.Printf("\r\n")

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Printf("before reverse: ")
	t = head
	for t != nil {
		fmt.Printf(" %d", t.Val)
		t = t.Next
	}
	fmt.Printf("\r\n")
	head = reverseKGroup(head, 3)
	fmt.Printf("after reverse: ")
	t = head
	for t != nil {
		fmt.Printf(" %d", t.Val)
		t = t.Next
	}
	fmt.Printf("\r\n")
}
