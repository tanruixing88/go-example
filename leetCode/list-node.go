package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func initListNode(nodeValList []int) *ListNode{
    var nxtNode *ListNode

    for i := len(nodeValList)-1; i >= 0; i-- {
        nxtNode = &ListNode{nodeValList[i], nxtNode}
    }

    return nxtNode
}

func printListNode(head *ListNode) {
    curNode := head
    fmt.Printf("Node List:")
    for curNode != nil {
        fmt.Printf("%d ", curNode.Val)
        curNode = curNode.Next
    }
    fmt.Printf("\r\n")
}

//25. K 个一组翻转链表
//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 || head == nil {
        return head
    }

    l := head
    h := l
    var r,ll *ListNode

    i := 1
    for h != nil {
        if i == k {
            tp := l
            t := tp.Next
            for t != h {
                tmp := t.Next
                t.Next = tp
                tp = t
                t = tmp
            }

            hn := h.Next
            t.Next = tp

            if ll == nil {
                r = t
            } else {
                ll.Next = t
            }
            ll = l

            l.Next = hn
            l = hn
            h = l
            i = 1
            continue
        }

        h = h.Next
        i++
    }

    return r
}

func main() {
	fmt.Printf("case1:\r\n")
	head := initListNode([]int{1,2,3,4,5})
	printListNode(head)
	r := reverseKGroup(head, 2)
    printListNode(r)
    fmt.Printf("case2:\r\n")
    head = initListNode([]int{1,2,3,4,5,6,7,8,9,10})
    printListNode(head)
    r = reverseKGroup(head, 3)
    printListNode(r)
}
