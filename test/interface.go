package main

import (
	"fmt"
	"reflect"
)

//输出为1,因为实际是append了[]int类型的元素
func appendInterface() {
	var nums1 []interface{}
	nums2 := []int{1,2,3}
	nums3 := append(nums1, nums2)
	fmt.Printf("%d\r\n", len(nums3))
}

//接口转换
type service interface {
	Do()
}

type v6 struct{}
func (*v6) Do() {}

func getV6() *v6 {
	return nil
}

func afterRefactor() {
	funcA(getV6())
}

func funcA(s service) {
	if s != nil {
		println("s is not nil. s:%+v\r\n", s)
	} else {
		println("s is nil")
	}
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}

	return false
}

//判断一个interface是否为nil
func interfaceIsNil() {
	var i interface{}
	IsNil(i)
}


//本质都是声明的int数组
func interfaceEqualArray() {
	var p [100]int
	var m interface{} = [...]int{99:0}
	fmt.Printf("interfaceEqualArray p == m:%t\r\n", p == m)
}


func main() {
	appendInterface()
	afterRefactor()
	interfaceIsNil()
	interfaceEqualArray()
}
