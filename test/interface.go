package main

import "fmt"

//输出为1,因为实际是append了[]int类型的元素
func appendInterface() {
	var nums1 []interface{}
	nums2 := []int{1,2,3}
	nums3 := append(nums1, nums2)
	fmt.Printf("%d\r\n", len(nums3))
}


func main() {
	appendInterface()
}
