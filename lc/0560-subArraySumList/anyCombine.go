package main

import "fmt"

//自由组合
func subArraySumList(nums []int, k int, rets *[][]int) [][]int {
	newLtKs := make([][]int, 0)

	if len(nums) == 0 {
		return newLtKs
	} else if len(nums) == 1 {
		if nums[0] < k {
			newLtKs = append(newLtKs, []int{nums[0]})
		} else if nums[0] == k {
			*rets = append(newLtKs, []int{nums[0]})
		}

		return newLtKs
	}

	tmpL := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		tmpL = append(tmpL, nums[i])
	}

	ltks := subArraySumList(tmpL, k, rets)

	for i := 0; i < len(ltks); i++ {
		newLtKs = append(newLtKs, ltks[i])
		sum := 0
		for j := 0; j < len(ltks[i]); j++ {
			sum += ltks[i][j]
		}

		if sum + nums[0] < k {
			tmpLtK := make([]int, 0)
			tmpLtK = append(tmpLtK, nums[0])
			for j := 0; j < len(ltks[i]); j++ {
				tmpLtK = append(tmpLtK, ltks[i][j])
			}
			newLtKs	= append(newLtKs, tmpLtK)
		} else if sum + nums[0] == k {
			tmpLtK := make([]int, 0)
			tmpLtK = append(tmpLtK, nums[0])
			for j := 0; j < len(ltks[i]); j++ {
				tmpLtK = append(tmpLtK, ltks[i][j])
			}
			*rets = append(*rets, tmpLtK)
		} else {
			continue
		}
	}

	if nums[0] == k {
		*rets = append(*rets, []int{nums[0]})
	} else if nums[0] < k {
		newLtKs = append(newLtKs, []int{nums[0]})
	}

	//fmt.Printf("nums:%+v newLtKs:%+v rets:%+v\r\n", nums, newLtKs, rets)

	return newLtKs
}

func anyCombineSum(nums []int, k int) int {
	rets := make([][]int, 0)
	subArraySumList(nums, k, &rets)
	//fmt.Printf("rets:%+v\r\n", rets)
	return len(rets)
}


func main() {
	nums := []int{1,1,1}
	k := 2
	fmt.Printf("subarraySum nums:%+v k:%d cnt:%d", nums, k, anyCombineSum(nums, k))
	nums = []int{1,2,3}
	k = 3
	fmt.Printf("subarraySum nums:%+v k:%d cnt:%d", nums, k, anyCombineSum(nums, k))
}
