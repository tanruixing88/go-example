package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(productList []int, queue chan int, close chan struct{}) {
	for _,product := range productList {
		queue <- product
	}

	close <- struct{}{}
}

func consumer(queue chan int, close chan struct{}) {
	for true {
		product, _ := <- queue
		fmt.Printf("%d ", product)
		if len(queue) == 0 {
			time.Sleep(time.Second)
			if len(queue) == 0 {
				break
			} else {
				continue
			}
		}
	}
	fmt.Printf("\r\n")
	close <- struct{}{}
}

func multiProducer(i int, wg *sync.WaitGroup, productList []int, queue chan int) {
	for _, product := range productList {
		queue <- product
	}

	wg.Done()
}

func multiConsumer(i int, wg *sync.WaitGroup, queue chan int, mutex *sync.Mutex) {
	for true {
		mutex.Lock()
		product, _ := <- queue
		fmt.Printf("%d ", product)
		if len(queue) == 0 {
			mutex.Unlock()
			time.Sleep(time.Second)
			if len(queue) == 0 {
				break
			} else {
				continue
			}
		}
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	var mutex sync.Mutex
	productList := []int{1,2,3,4,5,6,7,8,9}
	queue := make(chan int, 3)
	producerClose := make(chan struct{}, 1)
	go producer(productList, queue, producerClose)
	consumerClose := make(chan struct{}, 1)
	go consumer(queue, consumerClose)
	p := false
	c := false
	for true {
		select {
		case <-producerClose:
			fmt.Printf("producer close\r\n")
			p = true
		case <-consumerClose:
			fmt.Printf("consumer close\r\n")
			c = true
		}

		if p && c {
			break
		}
	}

	var producerWg sync.WaitGroup
	queue = make(chan int, 3)
	for i := 0; i < 3; i++ {
		initI := i * 10
		producerWg.Add(1)
		go multiProducer(i, &producerWg, []int{initI, initI+1, initI+2, initI+3, initI+4, initI+5}, queue)
	}

	var consumerWg sync.WaitGroup
	for j := 0; j < 3; j++ {
		consumerWg.Add(1)
		go multiConsumer(j, &consumerWg, queue, &mutex)
	}

	producerWg.Wait()
	consumerWg.Wait()
	fmt.Printf("\r\n")
	fmt.Printf("produers and consumers all done\r\n")
}
