package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMemStats(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s alloc memory:%vKB, gc count:%v\r\n", msg, m.Alloc/1024, m.NumGC)
}

//var m = make(map[int][]int)
func deleteMapKey() {
	m := make(map[int][]int)
	for i := 0; i < 100000; i++ {
		m[i] = make([]int, 1024, 1024)
	}
	printMemStats("add large keys")
	for i := 0; i < 100000; i++ {
		delete(m, i)
	}
	runtime.GC()
	printMemStats("del large keys")
	time.Sleep(1 * time.Second)
	runtime.GC() //加延迟时间并未有继续gc的处理, 内存和上次是一样的
	printMemStats("sleep del keys")


	m = nil //这样会整体去掉map
	runtime.GC()
	printMemStats("m = nil")
}

func deleteMapKeyCaller() {
	deleteMapKey() //函数结束后同样会回收
	runtime.GC()
	printMemStats("deleteMapKeyCaller1")
}


func main() {
	//golang map 删除key并不会释放内存
	deleteMapKeyCaller()
}
