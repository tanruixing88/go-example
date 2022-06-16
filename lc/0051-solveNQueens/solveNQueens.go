package main

import "fmt"

//此种解法和0039题有一定的共通之处, dfs 用法有类似之处
func solveNQueens(n int) [][]string {
	ret := make([][]string, 0)

	var dfs func(cols []int, row int)
	dfs = func(cols []int, row int) {
		if row == n {
			strs := make([]string, n)
			for j := 0; j < len(cols); j++ {
				str := make([]byte, n)
				for k := 0; k < n; k++ {
					if cols[j] == k {
						str[k] = 'Q'
					} else {
						str[k] = '.'
					}
				}
				strs[j] = string(str)
			}

			ret = append(ret, strs)
			return
		}

		abs := func(x int, y int) int { if x >= y {return x - y} else {return y - x}}
		for i := 0; i < n; i++ {
			valid := true
			for j := 0; j < len(cols); j++ {
				if i == cols[j] ||
					abs(row, j) == abs(i, cols[j]) {
					valid = false
					break
				}
			}

			if valid {
				cols = append(cols, i)
				dfs(cols, row + 1)
				cols = cols[:len(cols)-1]
			}
		}
	}

	dfs([]int{}, 0)

	return ret
}

func main() {
	n := 4
	fmt.Printf("n:%d queen:%+v\r\n", n, solveNQueens(n))

	n = 1
	fmt.Printf("n:%d queen:%+v\r\n", n, solveNQueens(n))
}
