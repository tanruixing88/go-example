package main

import (
	"fmt"
	"time"
)

/**
 * goroutine 本地队列大小256，
 *
 * 新加入G的情况, go func(){}()
 * 若本地队列不满，如果G是go新创建，或者目前的G是Gwaiter状态（time sleep，channel等触发），或者由于时间片（sysmon发现某个P执行某个G太久, 函数调用且栈内存大于128B）到达后，会放到队列后面
 * 若本地队列满，且其他P队列不满为空会被调度一半的G到空P中。
 * 若本地队列满，其他P队列满，则会调度一半的G到全局队列
 *
 * 完成G的情况
 * 若本地队列不空，则继续调度
 * 若本地队列为空，全局队列为空则会从其他P取一半的G,
 * 若本地队列为空，全局队列不空则从全局队列取当前P一半大小的G（或者全局队列G总数除P的数量的最小值）调度过来执行
 *
 * 若因为非阻塞系统调用则sysmon协程会调度解开P和M的关系，P继续和下一个空闲M（若没有则新建一个M）进行绑定并继续执行调度；将M和G关联进入内核态等待系统回调，如果回调完成M归入空闲队列等待新P继续绑定。
 *
 */




func main() {
	//main 函数结束都结束，非main函数，子协程都会存在
	fmt.Printf("main start\r\n")
	go func() {
		fmt.Printf("start \r\n")
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Printf("sub routine:%d\r\n", i)
				time.Sleep(time.Second)
			}
		}()
		fmt.Printf("end \r\n")
	}()
	fmt.Printf("main end\r\n")
	for true {
		time.Sleep(time.Second)
	}
}
