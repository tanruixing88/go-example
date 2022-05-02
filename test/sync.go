package main

import (
	"fmt"
	"sync"
)

func main() {
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
