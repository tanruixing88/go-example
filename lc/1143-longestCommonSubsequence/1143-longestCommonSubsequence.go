package main

import "fmt"

//https://leetcode-cn.com/problems/longest-common-subsequence/

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	if m == 0 || n == 0 {
		return 0
	}

	max := func(x, y int) int { if x > y {return x} else {return y}}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	r := longestCommonSubsequence("abcde", "ace")
	fmt.Printf("r:%d\r\n", r)
	r = longestCommonSubsequence("abc", "def")
	fmt.Printf("r:%d\r\n", r)
	r = longestCommonSubsequence("abc", "abc")
	fmt.Printf("r:%d\r\n", r)
	r = longestCommonSubsequence("hofubmnylkra","pqhgxgdofcvmr")
	fmt.Printf("r:%d\r\n", r) // except 5

}