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

func findLength1(nums1 []int, nums2 []int) int {
	maxL := 0
	for w := 1; w < len(nums1) + len(nums2); w++ {
		j := len(nums2)-w // len(nums2) - 1 - (w -1)
		if j < 0 {
			j = 0
		}

		i := w - len(nums2) // w - 1 - (len(nums2) - 1)
		if i < 0 {
			i = 0
		}

		t := 0
		tL := 0
		for i+t < len(nums1) && j+t < len(nums2) {
			if nums1[i+t] != nums2[j+t] {
				t++
				tL = 0
				continue
			}

			tL++
			if tL > maxL {
				maxL = tL
			}
			t++
		}
		fmt.Printf("w:%d i:%d j:%d tL:%d maxL:%d \r\n", w, i, j, tL, maxL)
	}

	return maxL
}

func main() {
	nums1 := []int{0,0,0,0,1}
	nums2 := []int{1,0,0,0,0}
	fmt.Printf("nums1:%+v nums2:%+v r:%d\r\n", nums1, nums2, findLength1(nums1, nums2))

	nums1 = []int{1,0,0,0,0}
	nums2 = []int{0,0,0,0,1}
	fmt.Printf("nums1:%+v nums2:%+v r:%d\r\n", nums1, nums2, findLength1(nums1, nums2))

	nums1 = []int{1,2,3}
	nums2 = []int{0,4,2,3,1}
	fmt.Printf("nums1:%+v nums2:%+v r:%d\r\n", nums1, nums2, findLength1(nums1, nums2))

	nums1 = []int{0,4,2,3,1}
	nums2 = []int{1,2,3}
	fmt.Printf("nums1:%+v nums2:%+v r:%d\r\n", nums1, nums2, findLength1(nums1, nums2))

	nums1 = []int{1,2,3,2,1}
	nums2 = []int{3,2,1,4,7}
	fmt.Printf("nums1:%+v nums2:%+v r:%d\r\n", nums1, nums2, findLength1(nums1, nums2))

	nums1 = []int{0,0,0,0,0}
	nums2 = []int{0,0,0,0,0}
	fmt.Printf("nums1:%+v nums2:%+v r:%d\r\n", nums1, nums2, findLength1(nums1, nums2))
}

