package main

import "fmt"

func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	preSum := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]

		if sumMap[preSum-k] > 0 {
			count += sumMap[preSum-k]
		}

		sumMap[preSum] += 1
	}

	return count
}


func main() {
	nums := []int{1,1,1}
	k := 2
	fmt.Printf("subarraySum nums:%+v k:%d cnt:%d\r\n", nums, k, subarraySum(nums, k))
	nums = []int{1,2,3}
	k = 3
	fmt.Printf("subarraySum nums:%+v k:%d cnt:%d\r\n", nums, k, subarraySum(nums, k))
	nums = []int{1}
	k = 0
	fmt.Printf("subarraySum nums:%+v k:%d cnt:%d\r\n", nums, k, subarraySum(nums, k))
	nums = []int{1,-1,0}
	k = 0
	fmt.Printf("subarraySum nums:%+v k:%d cnt:%d\r\n", nums, k, subarraySum(nums, k))
}
