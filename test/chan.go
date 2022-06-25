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
    fmt.Printf("simpleUse c1 is nil:%t, c2:%+v c3:%+v\r\n", &c1, c2, c3)
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
    for i := 0; i < 10; i++{
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
    after  := isChanClosed1(c1)

    fmt.Printf("testChanClosed before:%t after:%t\r\n", before, after)
}

func main() {
	simpleUse()
    randChan()
    forSelectOnlyOneCase()
	testChanClosed()

    chan2 := make(chan int)
    if chan2 == nil {
        fmt.Printf("no buffer chan is nil")
    } else {
        fmt.Printf("no buffer chan is not nil")
    }
    //直接进入休眠状态

    //chan3 := make(chan int, 10)
    //chan3 = nil
    //chan3 <- 1 //直接休眠
    //fmt.Printf("int value:%d\r\n", <-chan3) //直接休眠
    //fmt.Printf("can no reach\r\n")
}
