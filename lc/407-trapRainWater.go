package main

import "fmt"

type Point struct {
	X int
	Y int
}
// 效率并不高的解法
func trapRainWater(heightMap [][]int) int {
	trapNum := 0
	vMap := make(map[int]bool)
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			vMap[heightMap[i][j]] = true
		}
	}

	vList := make([]int, len(vMap))
	i := 0
	for k := range vMap {
		vList[i] = k
		i++
	}
	for i := 0; i < len(vList); i++ {
		for j := 0; j < len(vList) - 1 - i; j++ {
			if vList[j] > vList[j+1] {
				vList[j], vList[j+1] = vList[j+1], vList[j]
			}
		}
	}

	if len(vList) <= 1 {
		return 0
	}

	thMap := make([][]int, len(heightMap))
	for i := 0; i < len(heightMap); i++ {
		thMap[i] = make([]int, len(heightMap[i]))
	}
	for k := 0; k < len(vList)-1; k++ {
		v0Map := make(map[int]bool)
		for i := 0; i < len(heightMap); i++ {
			for j := 0; j < len(heightMap[i]); j++ {
				if heightMap[i][j] > vList[k] {
					thMap[i][j] = 1
				} else {
					if i == 0 || j == 0 || i == len(heightMap) - 1 || j == len(heightMap[i]) - 1 {
						thMap[i][j] = -1
					} else {
						thMap[i][j] = 0
						v0Map[i<<16+j] = true
					}
				}
			}
		}
		cnt := 1
		var del0Func func(int, int)
		del0Func = func(x int, y int) {
			if thMap[x+1][y] == 0 {
				thMap[x+1][y] = -1
				delete(v0Map, (x+1)<<16+y)
				del0Func(x+1, y)
			}

			if thMap[x-1][y] == 0 {
				thMap[x-1][y] = -1
				delete(v0Map, (x-1)<<16+y)
				del0Func(x-1, y)
			}

			if thMap[x][y+1] == 0 {
				thMap[x][y+1] = -1
				delete(v0Map, (x)<<16+y+1)
				del0Func(x, y+1)
			}

			if thMap[x][y-1] == 0 {
				thMap[x][y-1] = -1
				delete(v0Map, (x)<<16+y-1)
				del0Func(x, y-1)
			}
		}

		for true {
			change := false
			for k0 := range v0Map {
				ki := k0>>16
				kj := k0&((1<<16)-1)
				if thMap[ki+1][kj] == -1 || thMap[ki-1][kj] == -1 ||
					thMap[ki][kj+1] == -1 || thMap[ki][kj-1] == -1 {
					delete(v0Map, k0)
					change = true
					del0Func(ki, kj)
				}
			}

			if !change {
				break
			}
			cnt++
		}

		trapNum += (vList[k+1] - vList[k]) * len(v0Map)
	}

	return trapNum
}

//推荐的解法

func main() {
	r := trapRainWater([][]int{{1,4,3,1,3,2}, {3,2,1,3,2,4},{2,3,3,2,3,1}})
	fmt.Printf("r:%d\r\n", r) // except 4
	r = trapRainWater([][]int{{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}})
	fmt.Printf("r:%d\r\n", r) // except 10
	r = trapRainWater([][]int{{12,13,1,12},{13,4,13,12},{13,8,10,12},{12,13,12,12},{13,13,13,13}})
	fmt.Printf("r:%d\r\n", r) // except 14
}
