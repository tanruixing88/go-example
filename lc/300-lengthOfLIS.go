package main

import (
	"fmt"
)

func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	minL := 1
	for h := 1; h < len(nums); h++ {
		if nums[h-1] >= nums[h] {
			continue
		}

		tmpL := 2
		if tmpL > minL {
			minL = tmpL
		}

		tth := h
		pth := tth
		for tth < len(nums) {
			fmt.Printf("h:%d tth:%d \r\n", h, tth)
			f := false
			ttmpL := tmpL
			for th := tth + 1; th < len(nums); th++ {
				if nums[th] <= nums[pth] {
					fmt.Printf("h:%d tth:%d th:%d pth:%d\r\n", h, tth, th, pth)
					if !f && nums[th] > nums[h-1] {
						f = true
						tth = th
						tmpL++
						if tmpL > minL {
							minL = tmpL
						}
					}
					continue
				}
				pth = th
				ttmpL++
				if ttmpL > minL {
					minL = ttmpL
				}
			}

			if !f {
				break
			}
		}
	}

	return minL
}


func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dpMaxV := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dpMaxV < dp[j] {
					dpMaxV = dp[j]
				}
			}
		}
		dp[i] = dpMaxV + 1
	}

	dpMaxV := 0
	for i := 1; i < len(dp); i++ {
		if dpMaxV < dp[i] {
			dpMaxV = dp[i]
		}
	}

	return dpMaxV
}

func main() {
	r := lengthOfLIS([]int{10,9,2,5,3,7,101,18})
	fmt.Printf("r:%d \r\n", r)
	r = lengthOfLIS([]int{0,1})
	fmt.Printf("r:%d \r\n", r)
	r = lengthOfLIS([]int{0,1,0,3,2,3})
	fmt.Printf("r:%d \r\n", r)
	r = lengthOfLIS([]int{7,7,7,7,7,7})
	fmt.Printf("r:%d \r\n", r)
	r = lengthOfLIS([]int{10,9,2,5,3,4})
	fmt.Printf("r:%d \r\n", r)
}
