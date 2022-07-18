package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//slice 扩展是按照2倍扩展到1024,后续按照1.25倍进行扩展。

//通用删除某个元素的处理
func deleteElemByIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return []int{}
	}

	return append(s[0:index], s[index+1:]...)
}

func testDeleteElemByIndex() {
	s := []int{1,2,3,4}
	fmt.Printf("s:%+v", s)
	index := 0
	d := deleteElemByIndex(s, index)
	fmt.Printf("deleteElemByIndex index:%d s:%+v\r\n", index, d)

	s = []int{1,2,3,4}
	fmt.Printf("s:%+v", s)
	index = 2
	d = deleteElemByIndex(s, index)
	fmt.Printf("deleteElemByIndex index:%d s:%+v\r\n", index, d)

	s = []int{1,2,3,4}
	fmt.Printf("s:%+v", s)
	index = 3
	d = deleteElemByIndex(s, index)
	fmt.Printf("deleteElemByIndex index:%d s:%+v\r\n", index, d)
}


//空指针和空切片判断
func testNilAndEmptySlice() {
	var s1 []int
	s2 := make([]int, 0)
	s3 := make([]int, 0)

	//s1的data为0,但s2和s3的data值均为824634137496, nil的切片没有单独存储，但是所有空切片指向了同样的地址
	fmt.Printf("testNilAndEmptySlice s1 pointer:%+v, s2 pointer:%+v, s3 pointer:%+v\r\n",
	*(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)),
	*(*reflect.SliceHeader)(unsafe.Pointer(&s3)))

	//申请的长度不为0,则内存地址是不一样的
	s4 := make([]int, 1)
	s5 := make([]int, 1)
	fmt.Printf("testNilAndEmptySlice s4 pointer:%+v, s5 pointer:%+v\r\n",
		*(*reflect.SliceHeader)(unsafe.Pointer(&s4)), *(*reflect.SliceHeader)(unsafe.Pointer(&s5)))
}

//range 是否能够导致死循环
//答案是不能的，range本质就是先确定了最初s的长度，
// The loop we generate:
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//
func testForRangeLoop() {
	fmt.Printf("testForRangeLoop:")
	s := []int{1,2,3,4,5}
	for _, v := range s {
		s = append(s, v)
		//输出6 7 8 9 10
		fmt.Printf("len(s)=%d ", len(s))
	}
	fmt.Printf("\r\n")
}

func testAppend() {
	// s, a, b 地址都不一样
	s := []int{1,2,3}
	b := append(s, 4)
	a := append(s, 5)
	fmt.Printf("testAppend s:%+v a:%+v\r\n", s, a)
	fmt.Printf("testAppend s addr:%p b addr:%p a addr:%p\r\n", &s, &b, &a)

	s1 := make([]int, 5, 10)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	b1 := append(s1, 4)
	c1 := append(s1, 5)
	//a1 := append(s1, 5)
	//fmt.Printf("testAppend s1:%+v a1:%+v b1:%+v\r\n", s1, a1, b1)
	//fmt.Printf("testAppend s1 addr:%p a1 addr:%p b1 addr:%p\r\n", &s1, &a1, &b1)
	// b1 c1 结果全为[1 2 3 0 0 5], c1的append会修改b1的值
	fmt.Printf("testAppend s1:%+v b1:%+v c1:%+v\r\n", s1,  b1, c1)
	fmt.Printf("testAppend s1 addr:%p b1 addr:%p c1 addr:%p\r\n", &s1[0], &b1[0], &c1[0])
}


func Test1(s []int) {
	fmt.Printf("Test1 append before s0:%p s1:%p\r\n", &s[0], &s[1])
	s = append(s, 0)
	fmt.Printf("Test1 append after s0:%p s1:%p\r\n", &s[0], &s[1])
	for i := range s {
		s[i]++
	}
}

func testAppendV1() {
	s1 := []int{1,2}
	s2 := s1
	fmt.Printf("testAppendV1 append before s1[0]:%p s1[1]:%p\r\n", &s1[0], &s1[1])
	fmt.Printf("testAppendV1 append before s2[0]:%p s2[1]:%p\r\n", &s2[0], &s2[1])
	s2 = append(s2, 3)
	fmt.Printf("testAppendV1 append after s2:%p s2:%p\r\n", &s2[0], &s2[1])
	Test1(s1)
	Test1(s2)
	fmt.Printf("testAppendV1 s1:%+v s2:%+v\r\n", s1, s2)
}


func common() {
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
	//s5 = append(s5, 4) //数组无法append

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
		s[2] = 5 //单独修改还是结果被修改，但是前面加了append后就不一样了,是因为容量不足发生了深拷贝
	}
	appendModifyFunc(s10)
	fmt.Printf("s10:%+v\r\n", s10)

	s11 := []int{1,2,3}
	fmt.Printf("s11:%+v\r\n", s11[len(s11):])
	fmt.Printf("s11:%+v\r\n", s11[3:]) // 若打印s11[4:]则会引发panic
}


func main() {
	common()
	testDeleteElemByIndex()
	testNilAndEmptySlice()
	testForRangeLoop()
	testAppend()
	testAppendV1()
}
