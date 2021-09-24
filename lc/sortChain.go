package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func sortChain(node1 *Node, node2 *Node) *Node {
	var nodeSort *Node

	if node1 == nil && node2 == nil {
		return nil
	} else if node1 != nil && node2 == nil {
		return node1
	} else if node1 == nil && node2 != nil {
		return node2
	}

	if node1.val < node2.val {
		nodeSort = node1
	} else {
		nodeSort = node2
	}

	tNode := &Node{nil, 0}
	for node1 != nil && node2 != nil {
		if node1.val < node2.val {
			tNode.next = node1
			node1 = node1.next
		} else {
			tNode.next = node2
			node2 = node2.next
		}
		tNode = tNode.next
	}

	if node1 == nil {
		tNode.next = node2
	}
	if node2 == nil {
		tNode.next = node1
	}

	return nodeSort
}

func main() {
	// chain1  1->3
	node3 := &Node{nil, 3}
	node1 := &Node{node3, 1}

	// chain2  2->4->6
	node6 := &Node{nil, 6}
	node4 := &Node{node6, 4}
	node2 := &Node{node4, 2}

	node0 := sortChain(node1, node2)
	tmpNode := node0
	for tmpNode != nil {
		fmt.Printf(" %d ", tmpNode.val)
		tmpNode = tmpNode.next
	}
}
