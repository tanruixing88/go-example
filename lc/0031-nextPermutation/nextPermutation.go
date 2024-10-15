package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			minVal := nums[i]
			minIdx := i
			for j := i; j < n; j++ {
				if nums[j] > nums[i-1] && nums[j] < minVal {
					minVal = nums[j]
					minIdx = j
				}
			}
			nums[i-1], nums[minIdx] = nums[minIdx], nums[i-1]
			for j := minIdx; j < n-1; j++ {
				if nums[j] < nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
			//交换完毕后需要整体做一下升序处理,原先全部都是降序的，只需要反转一下即可
			// 2, 3, 1, 5, 4
			fmt.Printf("i:%d minIdx:%d \r\n", i, minIdx)
			for j := i; j < (n-i)/2+i; j++ {
				nums[j], nums[n-1-(j-i)] = nums[n-1-(j-i)], nums[j]
			}

			return
		}
	}

	for i := 0; i <= n/2-1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("nums:%+v ", nums)
	nextPermutation(nums)
	fmt.Printf("after nums:%+v\r\n", nums)

	nums = []int{3, 2, 1}
	fmt.Printf("nums:%+v ", nums)
	nextPermutation(nums)
	fmt.Printf("after nums:%+v\r\n", nums)

	nums = []int{4, 3, 2, 1}
	fmt.Printf("nums:%+v ", nums)
	nextPermutation(nums)
	fmt.Printf("after nums:%+v\r\n", nums)

	nums = []int{1, 1, 5}
	fmt.Printf("nums:%+v ", nums)
	nextPermutation(nums)
	fmt.Printf("after nums:%+v\r\n", nums)

	nums = []int{1, 3, 2}
	fmt.Printf("nums:%+v ", nums)
	nextPermutation(nums)
	fmt.Printf("after nums:%+v\r\n", nums)

	nums = []int{2, 3, 1, 3, 3}
	fmt.Printf("nums:%+v ", nums)
	nextPermutation(nums)
	fmt.Printf("after nums:%+v\r\n", nums)
}
