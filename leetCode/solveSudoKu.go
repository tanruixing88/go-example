package main

//官方的解决方案
// 见：https://leetcode-cn.com/problems/sudoku-solver/submissions/
func solveSudoku(board [][]byte)  {
	var rowMap, colMap [9][9]bool
	var subMap [3][3][9]bool
	var spaceList [][2]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				spaceList = append(spaceList, [2]int{i,j})
			} else {
				d := board[i][j] - '1'
				rowMap[i][d] = true
				colMap[j][d] = true
				subMap[i/3][j/3][d] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaceList) {
			return true
		}

		x := spaceList[pos][0]
		y := spaceList[pos][1]
		for k := 0; k < 9; k++ {
			if !rowMap[x][k] && !colMap[y][k] && !subMap[x/3][y/3][k] {
				rowMap[x][k] = true
				colMap[y][k] = true
				subMap[x/3][y/3][k] = true
				board[x][y] = byte(k) + '1'
				if dfs(pos+1) {
					return true
				} else {
					rowMap[x][k] = false
					colMap[y][k] = false
					subMap[x/3][y/3][k] = false
				}
			}
		}

		return false
	}

	dfs(0)
}

func main() {

}