package main

import "fmt"

//readNoDataFromNoBufChWithSelect 无阻塞无缓冲读
func readNoDataFromNoBufChWithSelect() {
	noBufCh := make(chan int)
	select {
	case x := <-noBufCh:
		fmt.Printf("read no buf channel. x:%d \r\n", x)
	default:
		fmt.Printf("readNoDataFromNoBufChWithSelect default\r\n")
	}
}

//readNoDataFromHasBufChWithSelect 无阻塞有缓冲读
func readNoDataFromHasBufChWithSelect() {
	bufCh := make(chan int, 1)
	select {
	case x := <-bufCh:
		fmt.Printf("read has buf channel. x:%d \r\n", x)
	default:
		fmt.Printf("readNoDataFromHasBufChWithSelect default\r\n")
	}
}

//writeDataToNoBufWithSelect 无阻塞无缓冲写
func writeDataToNoBufWithSelect() {
	noBufCh := make(chan int)
	select {
	case noBufCh <- 1:
		fmt.Printf("write to no buf channel success \r\n")
	default:
		fmt.Printf("writeDataToNoBufWithSelect default\r\n")
	}
}

//writeDataToHasBufWithSelect 无阻塞无缓冲写
func writeDataToHasBufWithSelect() {
	bufCh := make(chan int, 1)
	bufCh <- 2
	select {
	case bufCh <- 1:
		fmt.Printf("write to has buf channel success \r\n")
	default:
		fmt.Printf("writeDataToHasBufWithSelect default\r\n")
	}
}

func main() {
	readNoDataFromNoBufChWithSelect()
	readNoDataFromHasBufChWithSelect()
	writeDataToNoBufWithSelect()
	writeDataToHasBufWithSelect()
	chan1 := make(chan bool, 1)
	chan1 <- false
	select {
	case b1 := <-chan1:
		if b1 {
			fmt.Printf("true \r\n")
		} else {
			fmt.Printf("false \r\n")
		}
	}
}
