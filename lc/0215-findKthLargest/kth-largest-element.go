package main

import "fmt"

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

// topK的获取方式
func findKthLargest(nums []int, k int) int {
	//fmt.Printf("000000000 nums:%+v k:%d\r\n", nums, k)
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return -1
	}

	if k == 1 && len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if k == 1 {
			if nums[0] < nums[1] { return nums[1] } else { return nums[0] }
		} else {
			if nums[0] < nums[1] { return nums[0] } else { return nums[1] }
		}
	}

	p := nums[0]
	l := 1
	h := len(nums) - 1
	for l < h {
		if nums[l] <= p {
			l++
		} else if nums[h] >= p {
			h--
		} else {
			nums[l], nums[h] = nums[h], nums[l]
		}
	}

	//进行回退处理
	if nums[l] > p {
		l--
	} else if nums[h] < p {
		h++
	} else {
		l--
	}

	//fmt.Printf("l:%d h:%d\r\n", l, h)

	if h + k == len(nums) {
		min := nums[h]
		for i := h+1; i < len(nums); i++ {
			if nums[i] < min {
				min = nums[i]
			}
		}
		return min
	} else if h + k < len(nums) {
		return findKthLargest(nums[h:], k)
	} else {
		// p 比所有值都大的情况
		if h == len(nums) {
			if k - (len(nums) - h) == 1 {
				return p
			}
			return findKthLargest(nums[1:h], k - (len(nums) - h) - 1)
		}
		return findKthLargest(nums[:h], k - (len(nums) - h))
	}
}
// modify param num
func findKthLargestV2(nums []int, k int) int {
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

	nums = []int{-1,2,0}
	k = 2
	kth = findKthLargest(nums, k)
	fmt.Printf("nums:%+v k:%d kth:%d\r\n", nums, k, kth) //expect 1

	nums = []int{3,1,2,4}
	k = 2
	kth = findKthLargest(nums, k)
	fmt.Printf("nums:%+v k:%d kth:%d\r\n", nums, k, kth) //expect 1

	nums = []int{7,6,5,4,3,2,1}
	k = 2
	kth = findKthLargest(nums, k)
	fmt.Printf("nums:%+v k:%d kth:%d\r\n", nums, k, kth) //expect 1
}