package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	l := 1
	h := nums[0]
	for h <= len(nums) - 1 {
		maxH := h
		for i := l; i <= h; i++ {
			newH := nums[i] + i
			if newH >= len(nums) - 1 {
				return true
			}

			if newH > maxH {
				maxH = newH
			}
		}

		if maxH == h {
			return false
		}

		l = h+1
		h = maxH
	}

	return true
}

func main() {
	nums := []int{2,3,1,1,4}
	fmt.Printf("nums:%+v canJump:%t\r\n", nums, canJump(nums))

	nums = []int{3,2,1,0,4}
	fmt.Printf("nums:%+v canJump:%t\r\n", nums, canJump(nums))

}
