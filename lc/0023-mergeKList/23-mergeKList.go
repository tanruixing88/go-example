package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type AddrChain struct {
	Node *ListNode
	Next *AddrChain
}

func mergeKLists(lists []*ListNode) *ListNode {
	var addrH *AddrChain
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}

		if addrH == nil {
			addrH = &AddrChain{lists[i], nil}
		} else {
			cur := addrH
			if cur.Node.Val > lists[i].Val {
				addrH = &AddrChain{lists[i], cur}
				continue
			}

			for cur.Next != nil {
				if cur.Next.Node.Val > lists[i].Val {
					cur.Next = &AddrChain{lists[i], cur.Next}
					break
				}
				cur = cur.Next
			}

			if cur.Next == nil {
				cur.Next = &AddrChain{lists[i], nil}
			}
		}
	}

	var retNode *ListNode
	var curNode *ListNode

	if addrH != nil {
		retNode = addrH.Node
	} else {
		return nil
	}

	for addrH != nil {
		if addrH.Node.Next == nil {
			curNode = addrH.Node
			addrH = addrH.Next
			if addrH == nil {
				curNode.Next = nil
			} else {
				curNode.Next = addrH.Node
			}
			continue
		} else {
			curAddrNode := addrH

			for curAddrNode.Next != nil {
				if curAddrNode.Next.Node.Val > addrH.Node.Next.Val {
					curAddrNode.Next = &AddrChain{addrH.Node.Next, curAddrNode.Next}
					break
				}

				curAddrNode = curAddrNode.Next
			}

			if curAddrNode.Next == nil {
				curAddrNode.Next = &AddrChain{addrH.Node.Next, nil}
			}

			curNode = addrH.Node
			addrH = addrH.Next
			curNode.Next = addrH.Node
		}
	}

	return retNode
}

func main() {
	lists1 := []*ListNode{
		{1, &ListNode{4, &ListNode{5, nil}}},
		{1, &ListNode{3, &ListNode{4, nil}}},
		{2, &ListNode{6, nil}},
	}
	ret := mergeKLists(lists1)
	fmt.Printf("list1 merge ret:")
	i := 0
	for ret != nil && i < 100 {
		fmt.Printf("%d ", ret.Val)
		i++
		ret = ret.Next
	}
	fmt.Printf("\r\n")

	lists2 := []*ListNode{nil}
	ret = mergeKLists(lists2)
	fmt.Printf("list2 merge ret:")
	i = 0
	for ret != nil && i < 100 {
		fmt.Printf("%d ", ret.Val)
		i++
		ret = ret.Next
	}
	fmt.Printf("\r\n")

}
