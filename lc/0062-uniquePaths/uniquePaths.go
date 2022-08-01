package main

import "fmt"

//时间复杂度较高,不使用额外空间
func uniquePathsV1(m int, n int) int {
	pathNum := 0
	x := 0
	y := 0

	var dfs func(x int, y int)

	dfs = func(x int, y int) {
		if x >= m || y >= n {
			return
		}

		if x == m - 1 && y == n - 1 {
			pathNum++
			return
		}

		dfs(x+1, y)
		dfs(x, y+1)
	}

	dfs(x, y)
	return pathNum
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	m := 3
	n := 7
	pathNum := uniquePaths(m, n)
	fmt.Printf("m:%d n:%d pathNum:%d\r\n", m, n, pathNum)

	m = 3
	n = 2
	pathNum = uniquePaths(m, n)
	fmt.Printf("m:%d n:%d pathNum:%d\r\n", m, n, pathNum)

	m = 7
	n = 3
	pathNum = uniquePaths(m, n)
	fmt.Printf("m:%d n:%d pathNum:%d\r\n", m, n, pathNum)

	m = 3
	n = 3
	pathNum = uniquePaths(m, n)
	fmt.Printf("m:%d n:%d pathNum:%d\r\n", m, n, pathNum)

	m = 23
	n = 12
	pathNum = uniquePaths(m, n)
	fmt.Printf("m:%d n:%d pathNum:%d\r\n", m, n, pathNum)
}
