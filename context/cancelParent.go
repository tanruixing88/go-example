package main

import (
	"context"
	"fmt"
	"time"
)

func sleepRandom1(stopChan chan struct{}) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("This is sleep Random 1: %d\n", i)

		i++
		if i == 5 {
			fmt.Println("cancel sleep random 1")
			stopChan <- struct{}{}
			break
		}
	}
}

func sleepRandom2(ctx context.Context) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("This is sleep Random 2:%d\n", i)

		i++
		select {
			case <-ctx.Done():
				fmt.Printf("Why?%s\n", ctx.Err())
				fmt.Println("cancel sleep random 2")
				return
			default:
		}
	}
}

func main() {
	ctxParent, cancelParent := context.WithCancel(context.Background())
	ctxChild, _ := context.WithCancel(ctxParent)
	stopChan := make(chan struct{})
	go sleepRandom1(stopChan)
	go sleepRandom2(ctxChild)

	fmt.Printf("before select \r\n")

	select {
		case <- stopChan:
			fmt.Printf("stopChan received\r\n")
	}

	fmt.Printf("before cancel parent \r\n")

	cancelParent()
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Continue...\r\n")
	}
}
