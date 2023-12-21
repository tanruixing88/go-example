package main

import (
	"fmt"
	"sync"
)

func waitGroup() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		//必须是指针，否则就是值拷贝，陷入阻塞
		go func(wg *sync.WaitGroup, i int) {
			fmt.Printf("i:%d ", i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Printf("\r\n")
}

func syncMapRange() {
	var s sync.Map
	s.Store("1", "1")
	s.Store("2", "2")
	s.Store("3", "3")
	fmt.Printf("syncMapRange start range.\r\n")
	s.Range(func(kVal, vVal any) bool {
		key, _ := kVal.(string)
		val, _ := vVal.(string)
		fmt.Printf("key:%s val:%s\r\n", key, val)
		s.Delete("3")
		return false //这样会中断遍历
	})

	fmt.Printf("syncMapRange end range.\r\n")
	fmt.Printf("syncMapRange start get all k v.\r\n")
	s.Range(func(kVal, vVal any) bool {
		key, _ := kVal.(string)
		val, _ := vVal.(string)
		fmt.Printf("key:%s val:%s\r\n", key, val)
		return true
	})
	fmt.Printf("syncMapRange end get all k v.\r\n")
}

func main() {
	syncMapRange()
	waitGroup()
}
