package main

import (
	"fmt"
)

//回旋镖的数量
//447 https://leetcode-cn.com/problems/number-of-boomerangs/

func numberOfBoomerangs1(points [][]int) int {
	num := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				if (points[i][0] - points[j][0]) * (points[i][0] - points[j][0]) + (points[i][1] - points[j][1]) * (points[i][1] - points[j][1]) ==
					(points[i][0] - points[k][0]) * (points[i][0] - points[k][0]) + (points[i][1] - points[k][1]) * (points[i][1] - points[k][1]) {
					num++
				}

				if (points[j][0] - points[i][0]) * (points[j][0] - points[i][0]) + (points[j][1] - points[i][1]) * (points[j][1] - points[i][1]) ==
					(points[j][0] - points[k][0]) * (points[j][0] - points[k][0]) + (points[j][1] - points[k][1]) * (points[j][1] - points[k][1]) {
					num++
				}

				if (points[k][0] - points[j][0]) * (points[k][0] - points[j][0]) + (points[k][1] - points[j][1]) * (points[k][1] - points[j][1]) ==
					(points[k][0] - points[i][0]) * (points[k][0] - points[i][0]) + (points[k][1] - points[i][1]) * (points[k][1] - points[i][1]) {
					num++
				}
			}
		}
	}

	return num * 2
}

// 官方解法
func numberOfBoomerangs(points [][]int) int {
	num := 0
	for i := 0; i < len(points); i++ {
		dMap := make(map[int]int)
		for j := 0; j < len(points); j++ {
			s := (points[i][0] - points[j][0]) * (points[i][0] - points[j][0]) + (points[i][1] - points[j][1]) * (points[i][1] - points[j][1])
			dMap[s]++
		}

		for _, v := range dMap {
			num += v * (v - 1)
		}
	}

	return num
}

func main() {
	r := numberOfBoomerangs([][]int{{1,1}, {2,2}, {3,3}})
	fmt.Printf("r:%d\r\n", r)
}