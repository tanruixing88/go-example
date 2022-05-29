package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	h := len(height) - 1
	maxSum := 0

	for l < h {
		sum := 0
		if height[l] > height[h] {
			sum = (h - l) * height[h]
			h--
		} else {
			sum = (h - l) * height[l]
			l++
		}

		if maxSum < sum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	maxSum := maxArea(height)
	fmt.Printf("height:%+v maxSum:%d\r\n", height, maxSum)
	height = []int{1,1}
	maxSum = maxArea(height)
	fmt.Printf("height:%+v maxSum:%d\r\n", height, maxSum)

}
