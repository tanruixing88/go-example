package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	sumList := make([][]int, 0)
	for i := 0; i < len(nums) - 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums) - 2; j++ {
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}

			v := target - nums[i] - nums[j]

			l := j + 1
			h := len(nums) - 1
			for l < h {
				if h < len(nums) - 1 && nums[h] == nums[h+1] {
					h--
					continue
				}

				if l > j + 1 && nums[l] == nums[l-1] {
					l++
					continue
				}

				if nums[l] + nums[h] > v {
					h--
				} else if nums[l] + nums[h] < v {
					l++
				} else {
					sumList = append(sumList, []int{nums[i], nums[j], nums[l], nums[h]})
					l++
					h--
				}
			}
		}
	}

	return sumList
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	sumList := fourSum(nums, target)
	fmt.Printf("nums:%+v target:%d sumList:%+v \r\n", nums, target, sumList)

	nums = []int{2, 2, 2, 2, 2}
	target = 8
	sumList = fourSum(nums, target)
	fmt.Printf("nums:%+v target:%d sumList:%+v \r\n", nums, target, sumList)
}
