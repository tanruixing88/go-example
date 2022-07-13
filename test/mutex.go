package main

import (
	"fmt"
	"sync"
	"time"
)

//mutex 正常模式和饥饿模式以及mutex的使用注意事项
//https://blog.csdn.net/qq_37102984/article/details/115322706
//普通模式会goroutine不会进入等待队列，直接cpu层获取。若超过1ms则进入饥饿模式，
//新的goroutine直接进入等待队列尾部，而在队列头部的旧goroutine会被唤醒

func delayUnlock() {
	var m sync.Mutex
	fmt.Printf("A, ")
	m.Lock()
	go func() {
		time.Sleep(200 * time.Millisecond)
		m.Unlock()
		fmt.Printf("c ")
	}()
	m.Lock()
	fmt.Printf("B")
}

func main() {
	delayUnlock()
}
