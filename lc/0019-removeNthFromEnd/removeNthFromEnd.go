package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n <= 0 || head == nil {
        return nil
    }

    f := head
    for i := 0; i < n; i++ {
        if f == nil {
            return f
        }

        f = f.Next
    }

    var preE *ListNode
    e := head
    for f != nil {
        f = f.Next
        preE = e
        e = e.Next
    }

    if preE == nil {
        return head.Next
    } else {
        preE.Next = e.Next
        return head
    }
}

func printChain(head *ListNode) string {
    t := head
    outPut := ""
    for t != nil {
        outPut += fmt.Sprintf("%d ", t.Val)
        t = t.Next
    }

    return outPut
}

func main() {
    head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
    fmt.Printf("head:%s ", printChain(head))
    n := 2
    ret := removeNthFromEnd(head, n)
	fmt.Printf("n:%d ret:%s\r\n", n, printChain(ret))

    head = &ListNode{1, nil}
    fmt.Printf("head:%s ", printChain(head))
    n = 1
    ret = removeNthFromEnd(head, n)
    fmt.Printf("n:%d ret:%s\r\n", n, printChain(ret))

    head = &ListNode{1, &ListNode{2, nil}}
    fmt.Printf("head:%s ", printChain(head))
    n = 1
    ret = removeNthFromEnd(head, n)
    fmt.Printf("n:%d ret:%s\r\n", n, printChain(ret))
}
