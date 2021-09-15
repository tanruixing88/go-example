package main

import "fmt"

func main() {
	chan1 := make(chan bool, 1)
	chan1 <- false
	select {
		case b1 := <- chan1:
			if b1 {
				fmt.Printf("true \r\n")
			} else {
				fmt.Printf("false \r\n")
			}
	}
}
