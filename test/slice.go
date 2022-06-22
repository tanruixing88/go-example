package main

import (
	"fmt"
	"reflect"
)

//slice 扩展是按照2倍扩展到1024,后续按照1.25倍进行扩展。

func main() {
	var s1 [1]int
	s1[0] = 1
	//s1 长度是一的数组，不是slice类型, 不能用append
	var s2 []int
	s2 = append(s2, 1)
	//s1 = append(s2, 2) //ide 会提示报错
	fmt.Printf("s1 type:%s s2 type:%s \r\n", reflect.TypeOf(s1), reflect.TypeOf(s2))

	var s3 [10]int
	s4 := s3[5:6]
	fmt.Printf("s3 type:%s s4 type:%s \r\n", reflect.TypeOf(s3), reflect.TypeOf(s4))
	fmt.Printf("s3 len:%d cap:%d\r\n", len(s3), cap(s3))
	fmt.Printf("s4 len:%d cap:%d\r\n", len(s4), cap(s4))

	orderLen := 5
	order := make([]uint16, 2 * orderLen)
	for i := 0; i < len(order); i++ {
		order[i] = uint16(i + 1)
	}
	pollOrder := order[:orderLen:(orderLen+4)] // 这里的双冒号是指新的切片cap, 新的cap受制于原有slice cap长度限制
	lockOrder := order[orderLen:][:orderLen:orderLen]
	fmt.Printf("pollOrder len:%d cap:%d val:%+v\r\n", len(pollOrder), cap(pollOrder), pollOrder)
	fmt.Printf("lockOrder len:%d cap:%d val:%+v\r\n", len(lockOrder), cap(lockOrder), lockOrder)

	var s5 [3]int
	s5[0] = 1
	s5[1] = 2
	s5[2] = 3
	fmt.Printf("s5 slice 0 :%+v  s5 3:%+v len(s5):%d cap(s5):%d  s5 type:%s\r\n", s5[:0], s5[:3], len(s5), cap(s5), reflect.TypeOf(s5))
	//s5 = append(s5, 4)

	s6 := []int{1,2,3,4}
	s7 := append(s6[:1], s6[2:]...)
	fmt.Printf("delete elem 2. s6:%+v s7:%+v\r\n", s6, s7)

	//声明空和new的[]int均为nil
	var s8 []int
	s8IsNil := false
	if s8 == nil {
		s8IsNil = true
	}
	s9 := *new([]int)
	s9IsNil := false
	if s9 == nil {
		s9IsNil = true
	}

	fmt.Printf("s8:%t s9:%t\r\n", s8IsNil, s9IsNil)

	s10 := []int{1,2,3}
	appendModifyFunc := func(s []int) {
		s = append(s, 4) //  单纯append并不会改变切片里的值
		s[2] = 5 //单独修改还是结果被修改，但是前面加了append后就不一样了
	}
	appendModifyFunc(s10)
	fmt.Printf("s10:%+v\r\n", s10)

	s11 := []int{1,2,3}
	fmt.Printf("s11:%+v\r\n", s11[len(s11):])
	fmt.Printf("s11:%+v\r\n", s11[3:]) // 若打印s11[4:]则会引发panic

}
