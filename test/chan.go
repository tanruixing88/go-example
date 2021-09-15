package main

import (
    "fmt"
    "math/rand"
    "sync"
)

func randChan() {
    out := make(chan int)
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            out <- rand.Intn(5)
        }
        close(out)
    }()

    go func() {
        defer wg.Done()
        for i := range out {
            fmt.Printf("%d ", i)
        }
        fmt.Printf("\r\n")
    }()

    wg.Wait()
}

func main() {
    fmt.Println("hello world")

    chan1 := make(chan int, 10)
    chan1 <- 1
    fmt.Printf("int value:%d\r\n", <-chan1)

    fmt.Printf("chan1 len:%d\r\n", len(chan1))
    randChan()
}
