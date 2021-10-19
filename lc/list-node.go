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

func getListNodeCnt(head *ListNode) int {
    curNode := head
    cnt := 0
    for curNode != nil {
        curNode = curNode.Next
        cnt++
    }
    return cnt
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

// 142 环形链表 II
//https://leetcode-cn.com/problems/linked-list-cycle-ii/

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

	s := head
	f := head.Next

	for s != f {
	    if f == nil || f.Next == nil || f.Next.Next == nil {
	        return nil
        }

        f = f.Next.Next
        s = s.Next
    }

	pos := s.Next

	s = head
	for pos != s {
	    pos = pos.Next
	    s = s.Next
    }

	return pos
}

//链表排序 tNo:148
//https://leetcode-cn.com/problems/sort-list/

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    if head.Next.Next == nil {
        if head.Val > head.Next.Val {
            head.Val, head.Next.Val = head.Next.Val, head.Val
        }
        return head
    }

    p1 := head
    p2 := head.Next
    for p2 != nil && p2.Next != nil {
        p1 = p1.Next
        p2 = p2.Next.Next
    }

    p2h := p1.Next
    p1.Next = nil

    l1 := sortList(head)
    l2 := sortList(p2h)
    /*
    fmt.Printf("l1:")
    printListNode(l1)
    fmt.Printf("l2:")
    printListNode(l2)
    c1 := getListNodeCnt(l1)
    c2 := getListNodeCnt(l2)
     */
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 == nil && l2 != nil {
        return l2
    } else if l1 != nil && l2 == nil {
        return l1
    }

    var r *ListNode
    if l1.Val < l2.Val {
        r = l1
    } else {
        r = l2
    }

    for l1 != nil && l2 != nil {
    	if l1.Val < l2.Val {
    	    //一定要循环完毕，不然l1和l2两个列表里的值相同不好处理
            for l1.Next != nil && l1.Next.Val <= l2.Val {
                l1 = l1.Next
            }

            if l1.Next == nil {
                l1.Next = l2
                break
            } else {
            	t := l1.Next
                l1.Next = l2
                l1 = t
            }
        } else  {
            for l2.Next != nil && l2.Next.Val <= l1.Val {
                l2 = l2.Next
            }

            if l2.Next == nil {
                l2.Next = l1
                break
            } else {
                t := l2.Next
                l2.Next = l1
                l2 = t
            }
        }
    }

    /*
    c3 := getListNodeCnt(r)
    if c1 + c2 != c3 {
        fmt.Printf("***********************\r\n")
        fmt.Printf("rrrr:")
        printListNode(r)
    }
     */
    return r
}
//合并数组 tID-21
//https://leetcode-cn.com/problems/merge-two-sorted-lists/submissions/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    hNode := &ListNode{0, nil}
    tNode := hNode
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tNode.Next = l1
            l1 = l1.Next
        } else {
            tNode.Next = l2
            l2 = l2.Next
        }
        tNode = tNode.Next
    }

    if l1 == nil {
        tNode.Next = l2
    }
    if l2 == nil {
        tNode.Next = l1
    }

    return hNode.Next
}

//tID-234   此题是不改变head链表的解法, 已经ac过
// https://leetcode-cn.com/problems/aMhZSa/submissions/
// 快慢指针的写法要注意
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := slow
    nodeValList := make([]int, 0)
    for fast.Next != nil && fast.Next.Next != nil {
        nodeValList = append(nodeValList, slow.Val)
        slow = slow.Next
        fast = fast.Next.Next
    }

    next := slow.Next

    //奇数
    if fast.Next == nil {
        //next = next.Next //首节点已经算进去了
    } else if fast.Next.Next == nil { //偶数
        nodeValList = append(nodeValList, slow.Val)
    }

    for len(nodeValList) > 0 && next != nil {
        if next.Val != nodeValList[len(nodeValList)-1] {
            return false
        }

        next = next.Next
        nodeValList = nodeValList[:len(nodeValList)-1]
    }

    if next == nil && len(nodeValList) == 0 {
        return true
    }

    return false
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
    fmt.Printf("sortList:\r\n")
    head = initListNode([]int{4,2,1,3})
    printListNode(head)
    r = sortList(head)
    printListNode(r)
    head = initListNode([]int{-1,5,3,4,0})
    printListNode(head)
    r = sortList(head)
    printListNode(r)
    head = initListNode([]int{4,19,14,5,-3,1,8,5,11,15})
    printListNode(head)
    r = sortList(head)
    printListNode(r)
    head = initListNode([]int{-84,142,41,-17,-71,170,186,183,-21,-76,76,10,29,81,112,-39,-6,-43,58,41,111,33,69,97,-38,82,-44,-7,99,135,42,150,149,-21,-30,164,153,92,180,-61,99,-81,147,109,34,98,14,178,105,5,43,46,40,-37,23,16,123,-53,34,192,-73,94,39,96,115,88,-31,-96,106,131,64,189,-91,-34,-56,-22,105,104,22,-31,-43,90,96,65,-85,184,85,90,118,152,-31,161,22,104,-85,160,120,-31,144,115})
    printListNode(head)
    r = sortList(head)
    printListNode(r)
}
