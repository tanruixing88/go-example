package main

import "fmt"

//题号220
//链接：https://leetcode-cn.com/problems/contains-duplicate-iii/

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k < 1 {
		return false
	}

	i := 0
	j := 1

	fmt.Printf("*******************")
	for i <= len(nums) - 1 && j <= len(nums) - 1 {
		fmt.Printf("i:%d j:%d abs:%d \r\n", i, j, abs(nums[i], nums[j]))
		if abs(nums[i], nums[j]) <= t {
			return true
		}

		if j - i == k {
			i++
			j = i + 1
			continue
		}

		if j - i < k {
			if j == len(nums) - 1 {
				i++
				if i == j {
					break
				} else {
					j = i + 1
				}
			} else {
				j++
			}
			continue
		}
	}

	return false
}

func main() {
	r := containsNearbyAlmostDuplicate([]int{1,2,3,1}, 3, 0)
	fmt.Printf("r:%t\r\n", r)

	r = containsNearbyAlmostDuplicate([]int{1,0,1,1}, 1, 2)
	fmt.Printf("r:%t\r\n", r)

	r = containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9}, 2, 3)
	fmt.Printf("r:%t\r\n", r)

	r = containsNearbyAlmostDuplicate([]int{1,2}, 0, 1)
	fmt.Printf("r:%t\r\n", r)

	r = containsNearbyAlmostDuplicate([]int{-5,5,5,5,5,15}, 6, 6)
	fmt.Printf("r:%t\r\n", r)
}
