package main

import "fmt"

//718
// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/

func getMaxL(n1, n2 []int) int {
	maxL := 0
	for i := 0; i < len(n1); i++ {
		t := i
		tL := 0
		j := 0
		for j < len(n2) && t < len(n1) {
			if n1[t] != n2[j] {
				tL = 0
				t++
				j++
				continue
			}
			t++
			j++
			tL++
			if tL > maxL {
				maxL = tL
			}
		}
	}

	return maxL
}

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	l1 := getMaxL(nums1, nums2)
	l2 := getMaxL(nums2, nums1)
	if l1 > l2 {
		return l1
	} else {
		return l2
	}
}

func main() {
	r := findLength([]int{0,0,0,0,1}, []int{1,0,0,0,0})
	fmt.Printf("r:%d\r\n", r)

}

