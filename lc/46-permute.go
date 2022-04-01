package main

import "fmt"

func permute(nums []int) [][]int {
	retPermute := make([][]int, 0)

	if len(nums) <= 0 {
		return retPermute
	}

	curNum := nums[0]
	leftNums := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		leftNums = append(leftNums, nums[i])
	}

	if len(leftNums) > 0 {
		leftRetPermute := permute(leftNums)
		for j := 0; j < len(leftRetPermute); j++ {
			for k := 0; k < len(leftRetPermute[j]) + 1; k++ {
				tmpList := make([]int, 0)
				if k == 0 {
					tmpList = append(tmpList, curNum)
					for l := 0; l < len(leftRetPermute[j]); l++ {
						tmpList = append(tmpList, leftRetPermute[j][l])
					}
				} else if k > 0 && k < len(leftRetPermute[j]) {
					for l := 0; l < len(leftRetPermute[j]); l++ {
						if k == l {
							tmpList = append(tmpList, curNum)
						}
						tmpList = append(tmpList, leftRetPermute[j][l])
					}
				} else {
					for l := 0; l < len(leftRetPermute[j]); l++ {
						tmpList = append(tmpList, leftRetPermute[j][l])
					}
					tmpList = append(tmpList, curNum)
				}
				retPermute = append(retPermute, tmpList)
			}
		}
	} else {
		retPermute = append(retPermute, []int{curNum})
	}

	return retPermute
}

func main() {
	nums := []int{1}
	retPermute := permute(nums)
	fmt.Printf("retPermute:%+v \r\n", retPermute)

	nums = []int{1, 2}
	retPermute = permute(nums)
	fmt.Printf("retPermute:%+v \r\n", retPermute)

	nums = []int{1, 2, 3}
	retPermute = permute(nums)
	fmt.Printf("retPermute:%+v \r\n", retPermute)

	nums = []int{1, 2, 3, 4}
	retPermute = permute(nums)
	fmt.Printf("retPermute:%+v \r\n", retPermute)
}

