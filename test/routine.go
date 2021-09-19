package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

/**
 * goroutine 本地队列大小256，

 *
 *
 * 新加入G的情况, go func(){}()
 * 若本地队列不满，如果G是go新创建，或者目前的G是Gwaiter状态（time sleep，channel等触发），或者由于时间片（sysmon发现某个P执行某个G太久, 函数调用且栈内存大于128B）到达后，会放到队列后面
 * 若本地队列满，则会调度一半的G到全局队列
 *
 * 完成G的情况
 * 为了保证调度的公平性，每个工作线程每进行61次调度就需要优先从全局运行队列中获取goroutine出来运行
 * 若本地队列不空，则继续调度
 * 若本地队列为空，全局队列为空则会从其他P取一半的G,
 * 若本地队列为空，全局队列不空则从全局队列取当前P一半大小的G（或者全局队列G总数除P的数量的最小值）调度过来执行
 *
 * 若因为非阻塞系统调用则sysmon协程(20us-10ms)会调度解开P和M的关系，P继续和下一个空闲M（若没有则新建一个M）进行绑定并继续执行调度；将M和G关联进入内核态等待系统回调，如果回调完成M归入空闲队列等待新P继续绑定。
 *
 */

/*
 * https://github.com/lifei6671/interview-go/blob/master/base/go-scheduler.md
 * 首先来看对runnext的cas操作。只有跟_p_绑定的当前工作线程才会去修改runnext为一个非0值，其它线程只会把runnext的值从一个非0值修改为0值，
 * 然而跟_p_绑定的当前工作线程正在此处执行代码，所以在当前工作线程读取到值A之后，不可能有线程修改其值为B(0)之后再修改回A。
 * 再来看对runq的cas操作。当前工作线程操作的是_p_的本地队列，只有跟_p_绑定在一起的当前工作线程才会因为往该队列里面添加goroutine而去修改runqtail，
 * 而其它工作线程不会往该队列里面添加goroutine，也就不会去修改runqtail，它们只会修改runqhead，
 * 所以，当我们这个工作线程从runqhead读取到值A之后，其它工作线程也就不可能修改runqhead的值为B之后再第二次把它修改为值A
 *（因为runqtail在这段时间之内不可能被修改，runqhead的值也就无法越过runqtail再回绕到A值），也就是说，代码从逻辑上已经杜绝了引发ABA的条件。
 */


/**
 * 打印goroutine栈信息
 *
 */
func getGoRoutineId() int64 {
	b := make([]byte, 64)
	l := runtime.Stack(b, true)
	fmt.Printf("go routine stack:%s", string(b))
	b = b[:l]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseInt(string(b), 10, 64)
	return n
}


func main() {
	getGoRoutineId()
	/*
	//main 函数结束都结束，非main函数，子协程都会存在
	fmt.Printf("main start\r\n")
	go func() {
		fmt.Printf("start \r\n")
		go func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("sub routine:%d\r\n", i)
				time.Sleep(time.Second)
			}
		}()
		fmt.Printf("end \r\n")
	}()
	fmt.Printf("main end\r\n")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
	*/
}
