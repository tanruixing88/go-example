package main

import "fmt"

func rangeSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 0)
	for _, val := range slice1 {
		slice2 = append(slice2, val)
	}

	fmt.Printf("slice2:%+v\r\n", slice2)
}

func rangeAddr() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val
	}

	//输出最后一个，值为3, 若m的value是int类型，值复制，则为2
	fmt.Printf("rangeAddr %d m:%+v\r\n", *m[2], m)
}

func rangeKVAddr() {
	s := []int{0, 1, 2, 3, 4, 5}
	for k, v := range s {
		fmt.Printf("k:%d k addr:%p v:%d v addr:%p\r\n", k, &k, v, &v)
	}

}

func main() {
	rangeSlice()
	rangeKVAddr()
	rangeAddr()
}
