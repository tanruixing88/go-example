package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 *ListNode

    addNext := 0
    tl3 := l3
    for l1 != nil || l2 != nil {
        curNodeVal := addNext
        if l1 != nil {
            curNodeVal += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            curNodeVal += l2.Val
            l2 = l2.Next
        }

        if tl3 == nil {
            tl3 = &ListNode{curNodeVal % 10, nil}
            l3 = tl3
        } else {
            tl3.Next = &ListNode{curNodeVal % 10, nil}
            tl3 = tl3.Next
        }

        addNext = curNodeVal / 10
    }

    if addNext > 0 {
        tl3.Next = &ListNode{addNext, nil}
    }

    return l3
}

func printListNodeChain(l *ListNode) {
	t := l
    for t != nil {
        fmt.Printf("%d ", t.Val)
        t = t.Next
    }
    fmt.Printf("\r\n")
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
    l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
    l3 := addTwoNumbers(l1, l2)
    printListNodeChain(l3)

    l1 = &ListNode{0, nil}
    l2 = &ListNode{0, nil}
    l3 = addTwoNumbers(l1, l2)
    printListNodeChain(l3)

    l1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
    l2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
    l3 = addTwoNumbers(l1, l2)
    printListNodeChain(l3)
}
