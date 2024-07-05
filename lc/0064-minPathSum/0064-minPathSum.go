package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return -1
	}

	pathSumGrid := make([][]int, len(grid))
	for i := range pathSumGrid {
		pathSumGrid[i] = make([]int, len(grid[i]))
	}

	sumPath := 0
	for i := 0; i < len(grid); i++ {
		sumPath += grid[i][0]
		pathSumGrid[i][0] = sumPath
	}

	sumPath = 0
	for j := 0; j < len(grid[0]); j++ {
		sumPath += grid[0][j]
		pathSumGrid[0][j] = sumPath
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			pathSumGrid[i][j] = pathSumGrid[i][j-1]
			if pathSumGrid[i-1][j] < pathSumGrid[i][j] {
				pathSumGrid[i][j] = pathSumGrid[i-1][j]
			}

			pathSumGrid[i][j] += grid[i][j]
		}
	}

	return pathSumGrid[len(grid)-1][len(grid[0])-1]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	minPathSumV := minPathSum(grid)
	fmt.Printf("grid:%+v minPathSumV:%d", grid, minPathSumV)
}
