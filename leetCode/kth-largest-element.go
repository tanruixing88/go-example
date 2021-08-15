package main

import "fmt"

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/


// modify param num
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		maxI := i
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[maxI] {
				maxI = j
			}
		}

		t := nums[maxI]
		nums[maxI] = nums[i]
		nums[i] = t

		if i == k -1 {
			return t
		}
	}

	return -1
}

// no modify param nums
func findKthLargest_1(nums []int, k int) int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return -1
	}

	kthList := make([]int, k)
	//if nums[i] < 0, should init kthList[i] = -9999999
	for i := 0; i < len(nums); i++ {
		kI := 0
		for j := 0; j < len(kthList); j++ {
			if kthList[j] <= nums[i] {
				kI = j
				break
			}
		}

		if nums[i] < kthList[len(kthList)-1] {
			continue
		}

		for rI := len(kthList)-1; rI > kI ; rI-- {
			kthList[rI] = kthList[rI-1]
		}
		kthList[kI] = nums[i]
		//fmt.Printf("kthList:%+v\r\n", kthList)
	}

	return kthList[len(kthList)-1]
}

func main()  {
	nums := []int{3,2,1,5,6,4}
	k := 2
	kth := findKthLargest(nums, k)
	fmt.Printf("nums:%+v k:%d kth:%d\r\n", nums, k, kth) //expect 5

	nums = []int{3,2,3,1,2,4,5,5,6}
	k = 4
	kth = findKthLargest(nums, k)
	fmt.Printf("nums:%+v k:%d kth:%d\r\n", nums, k, kth) //expect 4

	nums = []int{2,1}
	k = 2
	kth = findKthLargest(nums, k)
	fmt.Printf("nums:%+v k:%d kth:%d\r\n", nums, k, kth) //expect 1
}