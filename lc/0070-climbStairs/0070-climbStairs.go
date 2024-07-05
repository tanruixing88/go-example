package main

import "fmt"

func climbStairs1(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs1(n-1) + climbStairs1(n-2)
}

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	pre1Num := 2
	pre2Num := 1

	for i := 3; i < n; i++ {
		num := pre1Num + pre2Num
		pre2Num = pre1Num
		pre1Num = num
	}

	return pre1Num + pre2Num
}

func main() {
	fmt.Printf("n:%d climbStairs:%d\r\n", 45, climbStairs(45))
}
