package main

import "fmt"

func main() {
    fmt.Println("hello world")

    chan1 := make(chan int, 10)
    chan1 <- 1
    fmt.Printf("int value:%d", <-chan1)

}