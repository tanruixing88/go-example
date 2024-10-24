package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	min := func(x int, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	a := make([][]int, m+1)
	for i := range a {
		a[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		a[i][0] = i
	}

	for j := 1; j <= n; j++ {
		a[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				a[i][j] = a[i-1][j-1]
			} else {
				//a[i-1][j-1]对应替换，操作+1  a[i-1][j]对应删除  a[i][j-1]对应追加
				a[i][j] = min(a[i-1][j-1], min(a[i-1][j], a[i][j-1])) + 1
			}
		}
	}

	return a[m][n]
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Printf("word1:%s word2:%s minDistance:%d\r\n", word1, word2, minDistance(word1, word2))

	word1 = "intention"
	word2 = "execution"
	fmt.Printf("word1:%s word2:%s minDistance:%d\r\n", word1, word2, minDistance(word1, word2))

	word1 = ""
	word2 = "a"
	fmt.Printf("word1:%s word2:%s minDistance:%d\r\n", word1, word2, minDistance(word1, word2))
}
