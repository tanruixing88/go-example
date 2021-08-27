package main

import "fmt"

type ChainNode struct {
	key int
	val int
	prev *ChainNode
	next *ChainNode
}

type LRUCache struct {
	KeyMap map[int] *ChainNode
	ChainHeader *ChainNode
	ChainTail *ChainNode
	Cap     int
	Len     int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*ChainNode), nil, nil, capacity, 0}
}


func (this *LRUCache) Get(key int) int {
	chainNode, ok := this.KeyMap[key]
	if !ok {
		return -1
	}

	//考虑首
	if this.KeyMap[key] == this.ChainHeader {
	} else if this.KeyMap[key] == this.ChainTail {
		this.ChainTail = this.ChainTail.prev
		this.ChainTail.next.prev = nil

		this.ChainHeader.prev = this.ChainTail.next
		this.ChainTail.next.next = this.ChainHeader
		this.ChainHeader = this.ChainTail.next

		this.ChainTail.next = nil
	} else {
		chainTmp := this.KeyMap[key]
		chainTmp.next.prev = chainTmp.prev
		chainTmp.prev.next = chainTmp.next

		this.ChainHeader.prev = chainTmp
		chainTmp.prev = nil
		chainTmp.next = this.ChainHeader
		this.ChainHeader = chainTmp
	}

	return chainNode.val
}


func (this *LRUCache) Put(key int, value int)  {
	newNode := &ChainNode{key, value, nil, nil}
	_, ok := this.KeyMap[key]
	if !ok {
		if this.Len == this.Cap {
			delKey := this.ChainTail.key
			this.ChainTail = this.ChainTail.prev
			this.ChainTail.next.prev = nil
			this.ChainTail.next = nil
			//free node
			delete(this.KeyMap, delKey)
			this.Len--
		}
		if this.ChainHeader != nil {
			this.ChainHeader.prev = newNode
		}
		newNode.next = this.ChainHeader
		this.ChainHeader = newNode

		if this.ChainTail == nil {
			this.ChainTail = newNode
		}
		this.Len++
		this.KeyMap[key] = newNode
		return
	}

	//考虑首
	if this.KeyMap[key] == this.ChainHeader {
		return
	} else if this.KeyMap[key] == this.ChainTail {
		this.ChainTail = this.ChainTail.prev
		this.ChainTail.next.prev = nil

		this.ChainHeader.prev = this.ChainTail.next
		this.ChainTail.next.next = this.ChainHeader
		this.ChainHeader = this.ChainTail.next

		this.ChainTail.next = nil
	} else {
		chainTmp := this.KeyMap[key]
		chainTmp.next.prev = chainTmp.prev
		chainTmp.prev.next = chainTmp.next

		this.ChainHeader.prev = chainTmp
		chainTmp.prev = nil
		chainTmp.next = this.ChainHeader
		this.ChainHeader = chainTmp
	}

	return
}

func (this *LRUCache) Print() {
	curNode := this.ChainHeader
	for curNode != this.ChainTail {
		fmt.Printf("%d ", curNode.val)
		curNode = curNode.next
	}
	fmt.Printf("%d\r\n", this.ChainTail.val)
}


func main() {
	var getRet int
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Print()
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Print()
	getRet = lRUCache.Get(1)    // 返回 1
	fmt.Printf("get ret:%d\r\n", getRet)
	lRUCache.Print()
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Print()
	getRet = lRUCache.Get(2)    // 返回 -1 (未找到)
	fmt.Printf("get ret:%d\r\n", getRet)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Print()
	getRet = lRUCache.Get(1)    // 返回 -1 (未找到)
	fmt.Printf("get ret:%d\r\n", getRet)
	getRet = lRUCache.Get(3)    // 返回 3
	fmt.Printf("get ret:%d\r\n", getRet)
	getRet = lRUCache.Get(4)    // 返回 4
	fmt.Printf("get ret:%d\r\n", getRet)
}
