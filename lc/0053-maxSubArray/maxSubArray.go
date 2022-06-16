package main

import (
	"fmt"
)

func maxSubArray1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	sum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			if sum < 0 {
				sum = nums[i]
			} else {
				sum += nums[i]
			}

			if sum > maxSum {
				maxSum = sum
			}
		} else {
			if sum < 0 {
				sum = nums[i]
			} else {
				if sum + nums[i] > 0 {
					sum = sum + nums[i]
				} else {
					sum = nums[i]
				}
			}

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func maxSubArray(nums []int) int {
	sum := nums[0]
	maxSum := sum
	for i := 1; i < len(nums); i++ {
		//其实之前的sum只要一直大于0就可以复用，本质只需要判断sum即可
		if sum >= 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	sum := maxSubArray(nums)
	fmt.Printf("nums:%+v sum:%d\r\n", nums, sum)

	nums = []int{1}
	sum = maxSubArray(nums)
	fmt.Printf("nums:%+v sum:%d\r\n", nums, sum)

	nums = []int{5,4,-1,7,8}
	sum = maxSubArray(nums)
	fmt.Printf("nums:%+v sum:%d\r\n", nums, sum)
}
