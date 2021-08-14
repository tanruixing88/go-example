package main

import "fmt"

type LRUCache struct {
	KeyMap map[int]int
	KeyList []int
	Cap     int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]int), make([]int, capacity), capacity}
}


func (this *LRUCache) Get(key int) int {
	val, ok := this.KeyMap[key]
	if !ok {
		return -1
	}

	return val
}


func (this *LRUCache) Put(key int, value int)  {
	_, ok := this.KeyMap[key]
	if !ok {
		for i := this.Cap - 1; i > 0 ; i-- {
			this.KeyList[i] = this.KeyList[i-1]
		}
		this.KeyList[0] = value
		this.KeyMap[key] = value
		return
	}
	this.KeyMap[key] = value

	pos := 0
	for i := 0; i < this.Cap; i++ {
		if key == this.KeyList[i] {
			pos = i
			break
		}
	}

	if pos > 0 {
		for i := pos; i > 0; i-- {
			this.KeyList[i] = this.KeyList[i-1]
		}
	}

	this.KeyList[0] = value
}

func (this *LRUCache) Print() {
	fmt.Printf("list:%+v map:%+v\r\n", this.KeyList, this.KeyMap)
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
