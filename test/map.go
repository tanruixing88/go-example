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

func lenValue() {
	//本质是0-9的key声明map，下面的0也包含在内
	//'a' 字节也可以算作非负整数，对应ascii码值为97
	m := [...]int{
		8: 1,
		9: 2,
	}
	fmt.Printf("before modify len value:%d\r\n", len(m))

	m[0] = 3
	fmt.Printf("after modify len value:%d\r\n", len(m))
}

//会报fatal error: concurrent map read and map write
//
func concurrentReadWriteMap() {
	m := make(map[int]int)
	go func() {
		for {
			_ = m[1]
		}
	}()
	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}

func modVal() {
	intMap := make(map[int]int)
	intMap[1] = 1
	intMap[2] = 2
	v2, _ := intMap[2]
	fmt.Printf("v2:%d\r\n", v2)
	v2 = 3
	fmt.Printf("intMap:%+v\r\n", intMap)
}

type structA struct {
	Count int
}

func getACount() {
	aMap := make(map[string]structA)
	aMap["a"] = structA{
		Count: 1,
	}

	bVal := aMap["b"]
	fmt.Printf("bVal:%d\r\n", bVal.Count)
}

func main() {
	getACount()
	modVal()
	//golang map 删除key并不会释放内存
	deleteMapKeyCaller()
	lenValue()
	//concurrentReadWriteMap()
}
