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

func binaryTreePaths(root *TreeNode) []string {
	pathStrList := []string{}

	if root == nil {
		return []string{}
	}

	curNode := root
	traceList := []*Trace{&Trace{curNode, false, false}}

	for true {
		if traceList[len(traceList)-1].TraceL && traceList[len(traceList)-1].TraceR {
			if curNode.Left == nil && curNode.Right == nil {
				pathStr := fmt.Sprintf("%d", traceList[0].Node.Val)
				for i := 1; i < len(traceList); i++ {
					pathStr += "->" + fmt.Sprintf("%d", traceList[i].Node.Val)
				}
				pathStrList = append(pathStrList, pathStr)
			}

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

	return pathStrList
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
	paths := binaryTreePaths(node1)
	fmt.Printf("tree paths:%+v\r\n", paths)
	clearTree(node1)


	//construct binaryTree
	/* only root node
			   1
	*/
	node1 = &TreeNode{1, nil, nil}
	paths = binaryTreePaths(node1)
	fmt.Printf("tree paths:%+v\r\n", paths)
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
	paths = binaryTreePaths(node1)
	fmt.Printf("tree paths:%+v\r\n", paths)
	clearTree(node1)

	//construct binaryTree
	/*
		   1
		 /   \
		2     3
	   / \     \
	  5   6     4
	*/
	node6 = &TreeNode{6, nil, nil}
	node5 = &TreeNode{5, nil, nil}
	node4 = &TreeNode{4, nil, nil}
	node3 = &TreeNode{3, nil, node4}
	node2 = &TreeNode{2, node5, node6}
	node1 = &TreeNode{1, node2, node3}
	paths = binaryTreePaths(node1)
	fmt.Printf("tree paths:%+v\r\n", paths)
	clearTree(node1)

	//construct binaryTree
	/*
			   1
			 /   \
			2     3
		   /       \
		  5         4
	*/
	node5 = &TreeNode{5, nil, nil}
	node4 = &TreeNode{4, nil, nil}
	node3 = &TreeNode{3, nil, node4}
	node2 = &TreeNode{2, node5, nil}
	node1 = &TreeNode{1, node2, node3}
	paths = binaryTreePaths(node1)
	fmt.Printf("tree paths:%+v\r\n", paths)
	clearTree(node1)

	//construct binaryTree
	/*
			   1
			 /   \
			2     3
		   / \     \
		  5   6     4
	     /
	    7
	*/
	node7 = &TreeNode{7, nil, nil}
	node6 = &TreeNode{6, nil, nil}
	node5 = &TreeNode{5, node7, nil}
	node4 = &TreeNode{4, nil, nil}
	node3 = &TreeNode{3, nil, node4}
	node2 = &TreeNode{2, node5, node6}
	node1 = &TreeNode{1, node2, node3}
	paths = binaryTreePaths(node1)
	fmt.Printf("tree paths:%+v\r\n", paths)
	clearTree(node1)
}
