package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3600; i++ {
		fmt.Printf("hello world %d\r\n", i)
		time.Sleep(1 * time.Second)
	}
}
