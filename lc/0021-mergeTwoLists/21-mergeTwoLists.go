package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func sortChain(node1 *Node, node2 *Node) *Node {
	if node1 == nil && node2 == nil {
		return nil
	} else if node1 != nil && node2 == nil {
		return node1
	} else if node1 == nil && node2 != nil {
		return node2
	}

	var nodeSort *Node
	if node1.val < node2.val {
		nodeSort = node1
	} else {
		nodeSort = node2
	}

	for node1 != nil && node2 != nil {
		if node1.val < node2.val {
			for node1.next != nil && node1.next.val <= node2.val {
				node1 = node1.next
			}

			if node1.next == nil {
				node1.next = node2
				break
			} else {
				tNode := node1.next
				node1.next = node2
				node1 = tNode
			}
		} else {
			for node2.next != nil && node2.next.val <= node1.val {
				node2 = node2.next
			}

			if node2.next == nil {
				node2.next = node1
				break
			} else {
				tNode := node2.next
				node2.next = node1
				node2 = tNode
			}
		}
	}

	return nodeSort
}

func sortChain1(node1 *Node, node2 *Node) *Node {
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
	chain1 := &Node{&Node{nil, 3}, 1}

	// chain2  2->4->6
	chain2 := &Node{&Node{&Node{nil, 6}, 2}, 2}

	ret := sortChain(chain1, chain2)
	tmpNode := ret
	for tmpNode != nil {
		fmt.Printf(" %d ", tmpNode.val)
		tmpNode = tmpNode.next
	}
	fmt.Printf("\r\n")

	// chain3  2->2->2
	chain3 := &Node{&Node{&Node{nil, 2}, 2}, 2}
	// chain4  2->4->6
	chain4 := &Node{&Node{&Node{nil, 6}, 2}, 2}

	ret = sortChain(chain3, chain4)
	tmpNode = ret
	for tmpNode != nil {
		fmt.Printf(" %d ", tmpNode.val)
		tmpNode = tmpNode.next
	}
}
