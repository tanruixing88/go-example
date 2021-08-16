package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Trace struct {
	Node *TreeNode
	TraceL bool
	TraceR bool
}

func maxPath(root *TreeNode, maxPathSum *int) int {

	maxValList := []int{}
	maxPathList := []int{}

	if root.Left == nil && root.Right == nil {
		if root.Val > *maxPathSum {
			*maxPathSum = root.Val
		}

		return root.Val
	} else if root.Left != nil && root.Right == nil {
		maxLeft := maxPath(root.Left, maxPathSum)
		maxValList = []int{root.Val, maxLeft+root.Val}
		maxPathList = []int{root.Val, maxLeft, maxLeft+root.Val}
	} else if root.Left == nil && root.Right != nil {
		maxRight := maxPath(root.Right, maxPathSum)
		maxValList = []int{root.Val, maxRight+root.Val}
		maxPathList = []int{root.Val, maxRight, maxRight+root.Val}
	} else {
		maxLeft := maxPath(root.Left, maxPathSum)
		maxRight := maxPath(root.Right, maxPathSum)
		maxValList = []int{root.Val, maxLeft + root.Val, maxRight + root.Val}
		maxPathList = []int{root.Val, maxLeft, maxLeft + root.Val,
			maxRight, maxRight + root.Val, maxLeft + root.Val + maxRight}
	}

	for i := 0; i < len(maxPathList); i++ {
		if *maxPathSum < maxPathList[i] {
			*maxPathSum = maxPathList[i]
		}
	}

	maxVal := maxValList[0]
	for i := 1; i < len(maxValList); i++ {
		if maxVal < maxValList[i] {
			maxVal = maxValList[i]
		}
	}

	return maxVal
}

// binary search tree path method
func maxPathSum1(root *TreeNode) int {
	maxPathSum := -999999

	if root == nil {
		return maxPathSum
	}


	maxPath(root, &maxPathSum)

	//fmt.Printf("root.Val:%d maxList:%+v maxPathSum:%d\r\n", root.Val, maxList, max)
	return maxPathSum
}

// binary search tree val method
func maxPathSum(root *TreeNode) int {
	maxPathSum := -999999

	max := func(x int, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode)	int {
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		nodePath := node.Val + leftGain + rightGain
		maxPathSum = max(nodePath, maxPathSum)

		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxPathSum
}


func clearTree(root *TreeNode) {
	if root == nil {
		return
	}

	curNode := root
	traceList := []*Trace{&Trace{curNode, false, false}}

	for true {
		if traceList[len(traceList)-1].TraceL && traceList[len(traceList)-1].TraceR {
			//fmt.Printf(" clear node:%+v\r\n", curNode)
			curNode = nil
			traceList = traceList[:len(traceList)-1]
			if len(traceList) == 0 {
				break
			}

			curNode = traceList[len(traceList)-1].Node
			continue
		}

		if curNode.Left != nil {
			if !traceList[len(traceList)-1].TraceL {
				traceList[len(traceList)-1].TraceL = true
				curNode = curNode.Left
				traceList = append(traceList, &Trace{curNode, false, false})
				continue
			}
		} else {
			traceList[len(traceList)-1].TraceL = true
		}

		if curNode.Right != nil {
			if !traceList[len(traceList)-1].TraceR {
				traceList[len(traceList)-1].TraceR = true
				curNode = curNode.Right
				traceList = append(traceList, &Trace{curNode, false, false})
				continue
			}
		} else {
			traceList[len(traceList)-1].TraceR = true
		}
	}
}

func main() {
	var node1,node2,node3,node4,node5,node6,node7 *TreeNode

	//construct binaryTree
	/* only root is nil
	 */
	max := maxPathSum(node1)
	fmt.Printf("max path sum:%+v\r\n", max)
	clearTree(node1)


	//construct binaryTree
	/* only root node
	   1
	*/
	node1 = &TreeNode{1, nil, nil}
	max = maxPathSum(node1)
	fmt.Printf("max path sum:%+v\r\n", max)
	clearTree(node1)

	//construct binaryTree
	/*

		   1
		 /   \
		2     3
		 \
		  5

	*/
	node5 = &TreeNode{5, nil, nil}
	node2 = &TreeNode{2, nil, node5}
	node3 = &TreeNode{3, nil, nil}
	node1 = &TreeNode{1, node2, node3}
	max = maxPathSum(node1)
	fmt.Printf("max path sum:%+v\r\n", max) // expect 11
	clearTree(node1)

	//construct binaryTree
	/*

		   1
		 /   \
		-2    -3
	   /  \     \
	  1	   3    -2
	 /
	-1

	*/
	node7 = &TreeNode{-1, nil, nil}
	node6 = &TreeNode{-2, node7, nil}
	node5 = &TreeNode{3, nil, nil}
	node4 = &TreeNode{1, nil, nil}
	node3 = &TreeNode{-3, nil, node6}
	node2 = &TreeNode{-2, node4, node5}
	node1 = &TreeNode{1, node2, node3}
	max = maxPathSum(node1)
	fmt.Printf("max path sum:%+v\r\n", max) // expect 3
	clearTree(node1)

	//construct binaryTree
	/*

			   9
			 /   \
			-3    0
		   /
		  3
		 /
		-8

	*/
	node5 = &TreeNode{-8, nil, nil}
	node4 = &TreeNode{3, node5, nil}
	node3 = &TreeNode{0, nil, nil}
	node2 = &TreeNode{-3, node4, nil}
	node1 = &TreeNode{9, node2, node3}
	max = maxPathSum(node1)
	fmt.Printf("max path sum:%+v\r\n", max) // expect 3
	clearTree(node1)
}
