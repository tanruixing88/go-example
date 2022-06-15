package main

import "fmt"

//交替打印1-100

func print1(chan1 chan bool, chan2 chan bool, chan3 chan bool) {
	for i := 1; i <= 100; i += 2 {
		<- chan1
		fmt.Printf("%d \r\n", i)
		chan2 <- true
	}

	chan3 <- true
}

func print2(chan1 chan bool, chan2 chan bool, chan3 chan bool) {
	for i := 2; i <= 100; i += 2 {
		<- chan2
		fmt.Printf("%d \r\n", i)
		chan1 <- true
	}
	chan3 <- true
}

func main() {
	chan1 := make(chan bool, 1)
	chan2 := make(chan bool, 1)
	chan3 := make(chan bool, 1)
	chan1 <- true
	go print1(chan1, chan2, chan3)
	go print2(chan1, chan2, chan3)
	<-chan3
	<-chan3
}
