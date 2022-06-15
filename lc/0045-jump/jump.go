package main

import "fmt"

func jump1(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	step := 1
	curIdxMap := make(map[int]bool)
	curIdxMap[len(nums)-1] = true
	for true {
		nxtIdxMap := make(map[int]bool)
		for k := range curIdxMap {
			for j := 0; j < k; j++ {
				if j + nums[j] >= k {
					if j == 0 {
						return step
					}
					//可能有重复的，重复赋值没有问题
					nxtIdxMap[j] = true
				}
			}
		}

		if len(nxtIdxMap) <= 0 {
			return step
		}
		curIdxMap = nxtIdxMap
		step++
	}

	return 0
}

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	l := 1
	h := nums[0]
	step := 1
	if h >= len(nums)-1 {
		return step
	}

	for true {
		maxH := 0
		for i := l; i <= h; i++ {
			curIdx := i + nums[i]
			if curIdx >= len(nums) - 1 {
				return step + 1
			}

			if curIdx > maxH {
				maxH = curIdx
			}
		}

		l = h + 1
		h = maxH
		step++
	}

	return step
}

func main() {
	nums := []int{2,3,1,1,4}
	step := jump(nums)
	fmt.Printf("nums:%+v jump:%d\r\n", nums, step)

	nums = []int{2,3,0,1,4}
	step = jump(nums)
	fmt.Printf("nums:%+v jump:%d\r\n", nums, step)
}
