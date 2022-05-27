package main

import "fmt"

func permute1(nums []int) [][]int {
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
		leftRetPermute := permute1(leftNums)
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

func permute( nums []int ) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	var recursion func(leftNums, tmpRes []int)

	recursion = func(leftNums, tmpRes []int) {
		if len(leftNums) == 0 {
			res = append(res, tmpRes)
			return
		}
		for index, value := range leftNums {
			newTmpRes := append(tmpRes, value)
			newLeftNums := make([]int, len(leftNums)-1)
			count := 0
			for i, v := range leftNums {
				if i != index {
					newLeftNums[count] = v
					count++
				}
			}
			// newLeftNums := append(leftNums[:index], leftNums[index+1:]...)
			// 不使用上面这行代码的原因是append会改变leftNums底层数组的值
			recursion(newLeftNums, newTmpRes)
		}
	}

	recursion(nums, tmp)
	return res

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

