package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	pathNumGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		pathNumGrid[i] = make([]int, n)
	}

	obstacleFlag := false
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] > 0 {
			obstacleFlag = true
		}

		if obstacleFlag {
			pathNumGrid[i][0] = 0
		} else {
			pathNumGrid[i][0] = 1
		}
	}

	obstacleFlag = false
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] > 0 {
			obstacleFlag = true
		}

		if obstacleFlag {
			pathNumGrid[0][j] = 0
		} else {
			pathNumGrid[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] > 0 {
				pathNumGrid[i][j] = 0
			} else {
				pathNumGrid[i][j] = pathNumGrid[i-1][j] + pathNumGrid[i][j-1]
			}
		}
	}

	/*
		fmt.Printf("pathNumGrid:\r\n")
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%d\t", pathNumGrid[i][j])
			}
			fmt.Printf("\r\n")
		}

	*/

	return pathNumGrid[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Printf("obstacleGrid:%+v pathNum:%d", obstacleGrid, uniquePathsWithObstacles(obstacleGrid))
}
