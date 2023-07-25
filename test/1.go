package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func buildTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}

	rootVal := preOrder[0]
	valIdx := 0
	for i := 0; i < len(inOrder); i++ {
		if inOrder[i] == rootVal {
			valIdx = i
			break
		}
	}

	return &TreeNode{rootVal, buildTree(preOrder[1:valIdx+1], inOrder[0:valIdx]), buildTree(preOrder[valIdx+1:], inOrder[valIdx+1:])}
}

func printNum() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			/*
				0 j 1 6 11
			    1 j
			*/
			for j := i + 1; j <= 100; j += 5 {
				fmt.Printf("i:%d Num:%d \r\n", i, j)
			}
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	fmt.Printf("3333333333333333333333 ")
}

func main() {
	printNum()
}
