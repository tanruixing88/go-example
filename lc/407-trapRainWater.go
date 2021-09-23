package main

import (
	"fmt"
)


// 效率并不高的解法
func trapRainWater1(heightMap [][]int) int {
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


func trapRainWater2(heightMap [][]int) int {
	fmt.Printf("start print:\r\n")
	for i := 0; i < len(heightMap); i++ {
		fmt.Printf("%+v\r\n", heightMap[i])
	}
	fmt.Printf("end print:\r\n")
	trapNum := 0
	M := len(heightMap)
	N := len(heightMap[0])
	if M < 3 || N < 3 {
		return 0
	}

	//处理i，j编码
	enHash := func(i, j int)int{return (i << 16) + j}
	deHash := func(i int)(int, int) {return i >> 16, i & ((1<<16)-1)}
	//堆处理
	addHeap := func(h []int, phI *int, e int) {
		h[*phI] = e
		*phI++
	}
	popHeap := func(h []int, shI, phI *int) int {
		//fmt.Printf("phI:%d shI:%d ", *phI, *shI)
		for i := *phI-1; i >= *shI; i--{
			j := i
			thI := (j-*shI-1)/2 + *shI
			//fmt.Printf(" j:%dthI:%d ", j, thI)
			for thI >= *shI {
				x1, y1 := deHash(h[j])
				x2, y2 := deHash(h[thI])
				if heightMap[x1][y1] < heightMap[x2][y2] {
					//fmt.Printf("swap:%d|%d ", j, thI)
					h[j], h[thI] = h[thI], h[j]
				}
				if thI == *shI {
					break
				}
				j = thI
				thI = (j-*shI-1)/2 + *shI
			}
		}
		*shI++
		return h[*shI-1]
	}

	hasMap := make(map[int]bool)
	hI := 0
	sI := 0
	var heap []int
	heap = make([]int, M * N + 4)
	for _, i := range []int{0, M-1} {
		for j := 0; j < N; j++ {
			addHeap(heap, &hI, enHash(i, j))
			hasMap[enHash(i, j)] = true
		}
	}
	for _, j := range []int{0, N-1} {
		for i := 0; i < M; i++ {
			addHeap(heap, &hI, enHash(i, j))
			hasMap[enHash(i, j)] = true
		}
	}

	inR := func(x, y int) bool { return x >= 1 && x < M-1 && y >= 1 && y < N-1}

	for true {
		v := popHeap(heap, &sI, &hI)
		if sI >= hI {
			break
		}
		cx, cy := deHash(v)
		for _, nv := range []int{enHash(cx,cy+1), enHash(cx,cy-1), enHash(cx+1,cy), enHash(cx-1,cy)} {
			nx, ny := deHash(nv)
			if inR(nx, ny) && !hasMap[enHash(nx, ny)] {
				hasMap[enHash(nx, ny)] = true
				if heightMap[nx][ny] < heightMap[cx][cy] {
					trapNum += heightMap[cx][cy] - heightMap[nx][ny]
					heightMap[nx][ny] = heightMap[cx][cy]
				}
				addHeap(heap, &hI, enHash(nx, ny))
			}
		}
	}

	return trapNum
}

//推荐的解法


type Point struct {
	X int
	Y int
}
func trapRainWater(heightMap [][]int) int {
	fmt.Printf("start print:\r\n")
	for i := 0; i < len(heightMap); i++ {
		fmt.Printf("%+v\r\n", heightMap[i])
	}
	fmt.Printf("end print:\r\n")
	trapNum := 0
	M := len(heightMap)
	N := len(heightMap[0])
	if M < 3 || N < 3 {
		return 0
	}

	//堆处理
	addHeap := func(h []Point, phI *int, e Point) {
		h[*phI] = e
		*phI++
	}
	popHeap := func(h []Point, shI, phI *int) Point {
		//fmt.Printf("phI:%d shI:%d ", *phI, *shI)
		minI := *shI
		for i := *shI; i < *phI; i++ {
			if heightMap[h[i].X][h[i].Y] < heightMap[h[minI].X][h[minI].Y] {
				minI = i
			}
		}
		h[*shI], h[minI] = h[minI], h[*shI]
		/*
		for i := *phI-1; i >= *shI; i--{
			j := i
			thI := (j-*shI-1)/2 + *shI
			//fmt.Printf(" j:%dthI:%d ", j, thI)
			for thI >= *shI {
				if heightMap[h[j].X][h[j].Y] < heightMap[h[thI].X][h[thI].Y] {
					//fmt.Printf("swap:%d|%d ", j, thI)
					h[j], h[thI] = h[thI], h[j]
				}
				if thI == *shI {
					break
				}
				j = thI
				thI = (j-*shI-1)/2 + *shI
			}
		}
		 */
		*shI++
		return h[*shI-1]
	}

	has := make([][]bool, M)
	for i := 0; i < M; i++ {
		has[i] = make([]bool, N)
	}

	hI := 0
	sI := 0
	heap := make([]Point, M * N + 4)
	for _, i := range []int{0, M-1} {
		for j := 0; j < N; j++ {
			addHeap(heap, &hI, Point{i, j})
			has[i][j] = true
		}
	}
	for _, j := range []int{0, N-1} {
		for i := 0; i < M; i++ {
			addHeap(heap, &hI, Point{i, j})
			has[i][j] = true
		}
	}

	inR := func(x, y int) bool { return x >= 1 && x < M-1 && y >= 1 && y < N-1}

	for true {
		v := popHeap(heap, &sI, &hI)
		if sI >= hI {
			break
		}
		cx := v.X
		cy := v.Y
		for _, nv := range []Point{{cx,cy+1}, {cx,cy-1}, {cx+1,cy}, {cx-1,cy}} {
			nx := nv.X
			ny := nv.Y
			if inR(nx, ny) && !has[nx][ny] {
				has[nx][ny] = true
				if heightMap[nx][ny] < heightMap[cx][cy] {
					trapNum += heightMap[cx][cy] - heightMap[nx][ny]
					heightMap[nx][ny] = heightMap[cx][cy]
				}
				addHeap(heap, &hI, Point{nx, ny})
			}
		}
	}

	return trapNum
}

func main() {
	r := trapRainWater([][]int{{78,16,94,36},{87,93,50,22},{63,28,91,60},{64,27,41,27},{73,37,12,69},{68,30,83,31},{63,24,68,36}})
	fmt.Printf("r:%d\r\n", r) // except 44
	r = trapRainWater([][]int{{5,8,7,7},{5,2,1,5},{7,1,7,1},{8,9,6,9},{9,8,9,9}})
	fmt.Printf("r:%d\r\n", r) // except 12
	r = trapRainWater([][]int{{1,4,3,1,3,2}, {3,2,1,3,2,4},{2,3,3,2,3,1}})
	fmt.Printf("r:%d\r\n", r) // except 4
	r = trapRainWater([][]int{{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}})
	fmt.Printf("r:%d\r\n", r) // except 10
	r = trapRainWater([][]int{{12,13,1,12},{13,4,13,12},{13,8,10,12},{12,13,12,12},{13,13,13,13}})
	fmt.Printf("r:%d\r\n", r) // except 14
	r = trapRainWater([][]int{{2,3,4},{5,6,7},{8,9,10},{11,12,13},{14,15,16}})
	fmt.Printf("r:%d\r\n", r) // except 0
	r = trapRainWater([][]int{{2,2,2},{2,1,2},{2,1,2},{2,1,2}})
	fmt.Printf("r:%d\r\n", r) // except 0
}
