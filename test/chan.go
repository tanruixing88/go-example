package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

func simpleUse() {
	chan1 := make(chan int, 10)
	chan1 <- 1
	fmt.Printf("simpleUse int value:%d\r\n", <-chan1)
	fmt.Printf("simpleUse chan1 len:%d\r\n", len(chan1))

	var c1 chan int
	c2 := make(chan int)
	c3 := make(chan int)
	//地址都是不一样的,值对应也不一样
	fmt.Printf("simpleUse c1 is nil:%p, c2:%+v c3:%+v\r\n", &c1, &c2, &c3)
}

//单通道的生产和消费
func randChan() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\r\n")
	}()

	wg.Wait()
}

//

const format = "2006-01-02 15:04:05"

func forSelectOnlyOneCase() {
	c := make(chan int)
	go func() {
		time.Sleep(3 * time.Millisecond)
		c <- 10
		close(c)
	}()

	//前几次输出都是走default分支，后续可以从c中读到数据10后，则走case分支
	//读到10,ok为true，剩余的循环取的x值为0,ok为false
	for i := 0; i < 10; i++ {
		select {
		case x, ok := <-c:
			fmt.Printf("forSelectOnlyOneCase %+v receive x:%d ok:%t\r\n", time.Now().Format(format), x, ok)
			time.Sleep(1 * time.Millisecond)
			if !ok {
				//c = nil // 若不去读一个已经关闭的chan，可以将chan赋值为nil, 或者c = make(chan int)
				c = make(chan int)
			}
		default:
			fmt.Printf("forSelectOnlyOneCase %+v receive nothing\r\n", time.Now().Format(format))
			time.Sleep(1 * time.Millisecond)
		}
	}
}

//判断一个chan是否关闭
func testChanClosed() {
	c1 := make(chan int16, 12)
	c1 <- 1
	c1 <- 2
	c1 <- 1
	c1 <- 6
	isChanClosed1 := func(c interface{}) bool {
		if reflect.TypeOf(c).Kind() != reflect.Chan {
			return false
		}

		//测试了hchan结构里的几个成员变量的值是符合预期的
		// this function will return true if chan.closed > 0
		// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
		// type hchan struct {
		// qcount   uint           // total data in the queue
		// dataqsiz uint           // size of the circular queue
		// buf      unsafe.Pointer // points to an array of dataqsiz elements
		// elemsize uint16
		// closed   uint32
		// **

		cptr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&c)) + unsafe.Sizeof(uint(0))))
		qcount := *(*uint)(unsafe.Pointer(cptr))
		cptr += unsafe.Sizeof(uint(0))
		dataqsiz := *(*uint)(unsafe.Pointer(cptr))
		fmt.Printf("qcount:%d dataqsize:%d\r\n", qcount, dataqsiz)
		fmt.Printf("array:%+v\r\n", unsafe.Sizeof(cptr))
		cptr += unsafe.Sizeof(uint(0))
		cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
		elemsize := *(*uint16)(unsafe.Pointer(cptr))
		fmt.Printf("elemSize:%d\r\n", elemsize)
		cptr += unsafe.Sizeof(uint16(0))
		fmt.Printf("isChanClosed1 val:%d\r\n", *(*uint32)(unsafe.Pointer(cptr)))
		return *(*uint32)(unsafe.Pointer(cptr)) > 0
	}

	before := isChanClosed1(c1)
	close(c1)
	after := isChanClosed1(c1)

	fmt.Printf("testChanClosed before:%t after:%t\r\n", before, after)
}

func emptyAndNilChan() {
	c1 := make(chan int)
	if c1 == nil {
		fmt.Printf("emptyAndNilChan no buffer chan is nil\r\n")
	} else {
		fmt.Printf("emptyAndNilChan no buffer chan is not nil\r\n")
	}

	// chan 为nil，是否阻塞, 只有select case: <- c,对应chan传递的阻塞变量为false
	/*
	   chan 为nil，则接收数据为非阻塞则直接返回，否则直接阻塞, 相关代码如下：
	   if c == nil {
	       if !block {
	           return
	       }
	       gopark(nil, nil, waitReasonChanReceiveNilChan, traceEvGoStop, 2)
	       throw("unreachable")
	   }

	   chan 为nil, 非阻塞则返回，阻塞则休眠
	     if c == nil {
	       if !block {
	         return false
	       }
	       gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
	       throw("unreachable")
	     }
	   总结 nil的chan发送和接收全是休眠，为nil且非阻塞情况还没找到对应的程序case
	*/
	go func() {
		//var c2 chan int
		c2 := make(chan int, 10)
		go func() {
			time.Sleep(100 * time.Millisecond)
			//c2 = nil //一旦被赋值则直接休眠
			a := <-c2
			fmt.Printf("emptyAndNilChan sleep c2 get a:%d\r\n", a)
		}()
		c2 <- 10
		//time.Sleep(200 * time.Millisecond)
		fmt.Printf("emptyAndNilChan c2 has send 10\r\n")
	}()
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("emptyAndNilChan nil chan can \r\n")
}

//输出为5
func readCloseChan() {
	c := make(chan int, 5)
	c <- 5
	c <- 6
	close(c)
	fmt.Printf("read c:%d\r\n", <-c)
}

func isChanPanic() {
	var c chan int
	//close(c) // close一个为nil的chan会panic。
	//v, ok := <-c // 直接从一个是nil的chan读取数据也会panic
	//c <- 11 // 给一个是nil的chan发送数据会panic

	c = make(chan int, 1)
	c <- 12
	close(c)
	//fmt.Printf("c is nil:%t \r\n", c == nil) //c不为nil
	//close(c) //重复close chan也会报错。
	//c <- 13 //给一个关闭的chan发送数据会panic
	//v, ok := <-c
	//fmt.Printf("isChanPanic v:%d ok:%t\r\n", v, ok)
}

func IsChanClose(c chan interface{}) bool {
	if c == nil {
		return true
	}

	_, ok := <-c
	return !ok
}

func closeChan() {
	c := make(chan interface{}, 1)
	close(c)
	fmt.Printf("chan c is closed:%t\r\n", IsChanClose(c))
}

func main() {
	closeChan()
	isChanPanic()
	simpleUse()
	randChan()
	forSelectOnlyOneCase()
	testChanClosed()
	emptyAndNilChan()
	readCloseChan()

	//直接进入休眠状态

	//chan3 := make(chan int, 10)
	//chan3 = nil
	//chan3 <- 1 //直接休眠
	//fmt.Printf("int value:%d\r\n", <-chan3) //直接休眠
	//fmt.Printf("can no reach\r\n")
}
