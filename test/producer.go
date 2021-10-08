package main
import (
	// ====== 可以引用Golang原生语言包 ====== //
	"fmt"
	"sync"
	"time"
)

// IWorkload 请勿修改接口
type IWorkload interface {
	// Work内包含一些耗时的处理，可能是密集计算或者外部IO
	Work()
}

// IProducer 请勿修改接口
type IProducer interface {
	// Produce每次调用会返回一个IWorkload实例
	// 当返回nil时表示已经生产完毕
	Produce() IWorkload
}

// 问题2：请编写函数Question2的实现如下功能
// 该函数输入一个IProducer实例，每次调用其Produce()方法会返回一个IWorkload实例。
// 1. 请反复调用该Produce()方法，直到返回nil，表明没有更多IWorkload。
//    此间可能会生产大量IWorkload实例，数目在此未知。
// 2. 对每个生产出的IWorkload实例，请调用一次它的Work()方法。
//    Work()内包含一些耗时的处理，可能是密集计算或者外部IO。
// 3. 请并发调用多个IWorkload的Work()方法，最多允许5个并发的Work()执行。
//    单个并发的实现，或并发数超过5的限制，都不能得分。
//
// 提示：请最小化内存、CPU代价
// 提示：请尽量使用规范的代码风格，使代码整洁易读
// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定

type Work2 struct {
	Val int
}

func (w Work2) Work() {
	t1 := time.Now()
	for true {
		if time.Now().Sub(t1).Seconds() > 1 {
			break
		}
	}
}

type Producer struct {
	Num *int
}

var cnt int = 0
func (p Producer)Produce() IWorkload {
	if *p.Num > 0 {
		*p.Num--
		return Work2{*p.Num}
	}

	cnt++
	return nil
}

func Question2(producer IProducer) {
	curNum := 5
	concurrentList := make([]chan IWorkload, curNum)
	for i := 0; i < curNum; i++ {
		concurrentList[i] = make(chan IWorkload, 2)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < curNum; i++ {
		wg.Add(1)
		go func(i int) {
			for w := range concurrentList[i] {
				if w == nil {
					fmt.Printf("produce done!\r\n")
					break
				}
				switch t := w.(type) {
				case Work2:
					fmt.Printf("goroutine num:%d work val:%d enter\r\n", i, t.Val)
					t.Work()
					fmt.Printf("goroutine num:%d work val:%d exit\r\n", i, t.Val)
				default:
					fmt.Printf("unkown\r\n")
				}
			}
			wg.Done()
		} (i)
	}
	go func(){
		endFlagList := make([]bool, curNum)
		for i := 0; true; i++ {
			idx := i % curNum
			allEnd := true
			for j := 0; j < curNum; j++ {
				if !endFlagList[j] {
					allEnd = false
					break
				}
			}
			if allEnd {
				break
			}

			if endFlagList[idx] {
				continue
			}

			workLoad := producer.Produce()
			switch w2 := workLoad.(type) {
			case Work2:
				fmt.Printf("w2 val:%d\r\n", w2.Val)
			}
			if workLoad == nil {
				fmt.Printf("workLoad is nil\r\n")
				endFlagList[idx] = true
			}
			concurrentList[idx] <- workLoad
			time.Sleep(1 * time.Millisecond)
		}

		for j := 0; j < curNum; j++ {
			close(concurrentList[j])
		}
		fmt.Printf("cnt:%d\r\n", cnt)
	}()
	wg.Wait()
}

func main() {
	a := 100
	Question2(Producer{&a})
}