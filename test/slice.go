package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s1 [1]int
	s1[0] = 1
	//s1 长度是一的数组，不是slice类型, 不能用append
	var s2 []int
	s2 = append(s2, 1)
	fmt.Printf("s1 type:%s s2 type:%s \r\n", reflect.TypeOf(s1), reflect.TypeOf(s2))

}
