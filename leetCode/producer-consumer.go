package main

import "fmt"

func producer(productList []int, queue chan int, close chan int) {
	for _,product := range productList {
		queue <- product
	}

	close <- 1
}

func consumer(queue chan int, close chan int) {
	for true {
		product, _ := <- queue
		fmt.Printf("%d ", product)
		if len(queue) == 0 {
			break
		}
	}
	fmt.Printf("\r\n")
	close <- 1
}

func main() {
	productList := []int{1,2,3,4}
	queue := make(chan int, 3)
	producerClose := make(chan int, 1)
	go producer(productList, queue, producerClose)
	consumerClose := make(chan int, 1)
	go consumer(queue, consumerClose)
	producerFlag, _ := <-producerClose
	consumerFlag, _ := <-consumerClose
	fmt.Printf("producer close flag:%d   consumer close flag:%d\r\n", producerFlag, consumerFlag)
}
