package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isTwoTreeOp(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }

    if root1 == nil && root2 != nil || root1 != nil && root2 == nil ||
        root1.Val != root2.Val {
       return false
    }

    return isTwoTreeOp(root1.Left, root2.Right) && isTwoTreeOp(root1.Right, root2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
	    return true
    }

	return isTwoTreeOp(root.Left, root.Right)
}

func printBinaryTree(root *TreeNode) {
	curNodeList := []*TreeNode{root}
	h := 0
	for len(curNodeList) > 0 {
		nxtNodeList := make([]*TreeNode, 0)
		for i := 0; i < len(curNodeList); i++ {
			if curNodeList[i] == nil {
				continue
			}

			if curNodeList[i].Left != nil {
				nxtNodeList = append(nxtNodeList, curNodeList[i].Left)
			}

			if curNodeList[i].Right != nil {
				nxtNodeList = append(nxtNodeList, curNodeList[i].Right)
			}
		}
		h++
		curNodeList = nxtNodeList
	}

	if h <= 0 {
		return
	}
	//以最底层的节点为长度基准
	//完全二叉树底层节点是 1<<(h-1), 考虑数值为十位数 还需要1<<(h-1)
	//对应分支也需要占据空间，一个节点对应一个分支 还需要 1<<(h-1)
	//一对左右孩子还需要一定的空隙，再加 (1<<(h-1)) / 2 * 2
	//共计 1<<(h-1) * 4 等于 1 <<(h+1)
	strLen := 1 << uint(h+1)

	curNodeList = []*TreeNode{root}
	curNodePosList := []int{strLen / 2}

	curH := 0
	for len(curNodeList) > 0 {
		branchDistance := 0
		if h-curH > 2 {
			branchDistance = 1 << uint(h-curH-3)
		}

		curStr := make([]byte, 0)
		for j := 0; j < strLen; j++ {
			curStr = append(curStr, ' ')
		}
		curBranchStr := make([]byte, 0)
		for j := 0; j < strLen; j++ {
			curBranchStr = append(curBranchStr, ' ')
		}

		nxtNodeList := make([]*TreeNode, 0)
		nxtNodePosList := make([]int, 0)

		for i := 0; i < len(curNodeList); i++ {
			if curNodeList[i] == nil {
				continue
			}

			tVal :=  curNodeList[i].Val
			curPos := curNodePosList[i]
			for true {
				curStr[curPos] = byte((tVal % 10) + '0')
				tVal = tVal / 10
				if tVal == 0 {
					break
				}
				curPos++
			}

			if curNodeList[i].Left != nil {
				nxtNodeList = append(nxtNodeList, curNodeList[i].Left)
				nxtNodePosList = append(nxtNodePosList, curPos-2-branchDistance)
				curBranchStr[curPos-1-branchDistance] = '/'
			}

			if curNodeList[i].Right != nil {
				nxtNodeList = append(nxtNodeList, curNodeList[i].Right)
				nxtNodePosList = append(nxtNodePosList, curPos+2+branchDistance)
				curBranchStr[curPos+1+branchDistance] = '\\'
			}
		}
		fmt.Printf("%s\r\n", curStr)
		fmt.Printf("%s\r\n", curBranchStr)
		curNodeList = nxtNodeList
		curNodePosList = nxtNodePosList
		curH++
	}
}

func main() {
	n7 := &TreeNode{7, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n3 := &TreeNode{3, n6, n7}
	n2 := &TreeNode{2, n4, n5}
	n1 := &TreeNode{1, n2, n3}
	printBinaryTree(n1)
	fmt.Printf("isSymmetric:%t\r\n", isSymmetric(n1))


}
