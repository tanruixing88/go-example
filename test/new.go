package main

import "fmt"

type HasMap struct {
	IntVal int
	IntMap map[int]int
}

func newStruct() {
	hasMap := new(HasMap)
	//hasMap.IntMap[1] = 1 //代码会崩溃，提示panic: assignment to entry in nil map
	fmt.Printf("new struct map:%+v\r\n", hasMap)
}

func main() {
	newStruct()

}
