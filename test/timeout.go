package main

import (
	"fmt"
	"time"
	// ====== 可以引用Golang原生语言包 ====== //
)

// IWorkload 请勿修改接口
type IWorkload interface {
	// Process内包含一些耗时的处理，可能是密集计算或者外部IO
	Process()
}

var TimeoutError = fmt.Errorf("timeout")

// 问题1：请编写函数Question1的实现如下功能
// 该函数输入一个IWorkload实例，请调用其Process函数一次，
// 调用完毕则Question1返回，此时返回的error应为空
// 当Process函数执行5秒仍未能结束时，让Question1函数不再等待
// 立即返回TimeoutError
//
// 注意：题目要求只调用Process一次
// 注意：超时时间固定5秒，请不要修改Question1函数的输入、输出定义
// 提示：请尽量使用规范的代码风格，使代码整洁易读
// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定

const SyncChanOpen = 1
const SyncChanClose = 2

func Question1(workload IWorkload) (err error) {
	syncChan := make(chan int)
	go func() {
		syncChan <- SyncChanOpen
		workload.Process()
		syncChan <- SyncChanClose
	}()
	//先让启动前统一唤醒主协程和子协程同时处理
	syncChanVal := <- syncChan
	timerChan:= time.After(5 * time.Second)
	fmt.Printf("start syncChanVal:%d\r\n", syncChanVal)

	for true {
		select {
			case <-timerChan:
				return TimeoutError
			case syncChanVal = <-syncChan:
				fmt.Printf("end syncChanVal:%d\r\n", syncChanVal)
				return nil
		}
	}

	return
}

type Work1  struct {
	Val int
}

func (t Work1) Process() {
	t1 := time.Now()
	for true {
		t.Val++
		//cpu 型耗时 5s  也可以直接修改为3s和6s进行测试
		if time.Now().Sub(t1).Seconds() > 6 {
			break
		}
	}
}

func main() {
	t :=  Work1{}
	err := Question1(t)
	if err == nil {
		fmt.Printf("err is nil\r\n")
		return
	}
	fmt.Printf("err:%s\r\n", err)
}
