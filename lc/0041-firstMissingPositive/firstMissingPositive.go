package main

import (
	"fmt"
	"math"
)

func firstMissingPositive(nums []int) int {
	if len(nums) < 1 {
		return -1
	}

	minPositive := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if minPositive > nums[i] && nums[i] > 0 {
			minPositive = nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("minPositive:%d nums:%+v i:%d\r\n", minPositive, nums, i)
		if nums[i] > 0 {
			curIdx := i
			nxtIdx := nums[curIdx] - minPositive
			for true {
				//fmt.Printf("1111 minPositive:%d nums:%+v i:%d curIdx:%d nxtIdx:%d\r\n", minPositive, nums, i, curIdx, nxtIdx)
				if nxtIdx >= len(nums) {
					break
				} else if curIdx == nxtIdx {
					break
				} else if nums[curIdx] == nums[nxtIdx] {
					break
				} else {
					nxtVal := nums[nxtIdx]
					nums[nxtIdx] = nums[curIdx]
					nums[curIdx] = nxtVal
					if nxtVal <= 0 {
						break
					} else {
						nxtIdx = nxtVal - minPositive
					}
				}
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		num := i + 1
		if nums[i] != num {
			return num
		}
	}

	return len(nums) + 1
}

func main() {
	nums := []int{1,2,0}
	next := firstMissingPositive(nums)
	fmt.Printf("nums:%+v next:%d \r\n", nums, next)

	nums = []int{3,4,-1,1}
	next = firstMissingPositive(nums)
	fmt.Printf("nums:%+v next:%d \r\n", nums, next)

	nums = []int{7,8,9,11,12}
	next = firstMissingPositive(nums)
	fmt.Printf("nums:%+v next:%d \r\n", nums, next)

	nums = []int{0}
	next = firstMissingPositive(nums)
	fmt.Printf("nums:%+v next:%d \r\n", nums, next)

	nums = []int{1}
	next = firstMissingPositive(nums)
	fmt.Printf("nums:%+v next:%d \r\n", nums, next)

	nums = []int{1,1}
	next = firstMissingPositive(nums)
	fmt.Printf("nums:%+v next:%d \r\n", nums, next)
}
