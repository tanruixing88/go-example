package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    t := head
    if t == nil {
    	return head
	}

    n := t.Next
    for t != nil && t.Next != nil {
       t.Val, n.Val = n.Val, t.Val
       if n.Next == nil || n.Next.Next == nil {
           return head
       }
       t = t.Next.Next
       n = n.Next.Next
    }

    return head
}

func printChain(node *ListNode) {
    for node != nil {
        fmt.Printf("%d ", node.Val)
        node = node.Next
    }
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Printf("head:")
	printChain(head)
	swapPairs(head)
	fmt.Printf(" swap:")
	printChain(head)
	fmt.Printf(" \r\n")


	head = nil
	fmt.Printf("head:")
	printChain(head)
	swapPairs(head)
	fmt.Printf(" swap:")
	printChain(head)
	fmt.Printf(" \r\n")

	head = &ListNode{1, nil}
	fmt.Printf("head:")
	printChain(head)
	swapPairs(head)
	fmt.Printf(" swap:")
	printChain(head)
	fmt.Printf(" \r\n")

}
