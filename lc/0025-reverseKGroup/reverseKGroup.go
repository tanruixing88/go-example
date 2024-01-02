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

	var pStartNode *ListNode
	var preNode *ListNode

	startNode := head
	curCount := 1
	endNode := startNode

	for endNode != nil {
		fmt.Printf("startNode:%d endNode:%d curCount:%d\r\n", startNode.Val, endNode.Val, curCount)
		if curCount < k-1 {
			curCount++
			endNode = endNode.Next
			fmt.Printf("1111111111 startNode:%d curCount:%d  endNode:%d\r\n", startNode.Val, curCount, endNode.Val)
		} else if curCount == k-1 {
			pStartNode = preNode
			preNode = endNode
			curCount++
			endNode = endNode.Next
			if endNode != nil {
				fmt.Printf("2222222 startNode:%d curCount:%d  endNode:%d\r\n", startNode.Val, curCount, endNode.Val)
			}
		} else {
			lNode := startNode
			rNode := endNode
			fmt.Printf("3333333 startNode:%d curCount:%d  endNode:%d\r\n", startNode.Val, curCount, endNode.Val)
			endNode = endNode.Next
			curCount = 1
			startNode = endNode

			cNode := lNode
			nNode := cNode.Next
			for nNode != rNode {
				fmt.Printf("44444 cNode:%d nNode:%d rNode:%d\r\n", cNode.Val, nNode.Val, rNode.Val)
				nnNode := nNode.Next
				nNode.Next = cNode
				cNode = nNode
				nNode = nnNode
			}

			lNode.Next = rNode.Next
			rNode.Next = cNode
			if pStartNode == nil {
				head = rNode
			} else {
				pStartNode.Next = rNode
			}
		}
	}

	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Printf("before reverse: ")
	t := head
	for t != nil {
		fmt.Printf(" %d", t.Val)
		t = t.Next
	}
	fmt.Printf("\r\n")
	head = reverseKGroup(head, 2)
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
