package main

import "fmt"

func minReorder(n int, connections [][]int) int {
	connTree := make([][][]int, n)
	for _, conn := range connections {
		connTree[conn[0]] = append(connTree[conn[0]], []int{conn[1], 1})
		connTree[conn[1]] = append(connTree[conn[1]], []int{conn[0], 0})
	}

	var dfs func(int, int, [][][]int) int
	dfs = func(curCity int, preCity int, connTree [][][]int) int {
		reOrderNum := 0
		for _, connCity := range connTree[curCity] {
			//更加合理的情况就是用一个map来解决之前访问过的节点
			if connCity[0] == preCity {
				continue
			}

			reOrderNum += dfs(connCity[0], curCity, connTree) + connCity[1]
		}

		return reOrderNum
	}

	return dfs(0, -1, connTree)
}

func minReorder_v3(n int, connections [][]int) int {
	reOrderNum := 0
	curConnList := connections
	hasTraverse := map[int]bool{0: true}

	for len(curConnList) > 0 {
		nxtConnList := make([][]int, 0)
		for _, curConn := range curConnList {
			if hasTraverse[curConn[1]] {
				hasTraverse[curConn[0]] = true
			} else if hasTraverse[curConn[0]] {
				hasTraverse[curConn[1]] = true
				reOrderNum++
			} else {
				nxtConnList = append(nxtConnList, curConn)
			}
		}

		curConnList = nxtConnList
	}

	if len(hasTraverse) == n {
		return reOrderNum
	}

	return 0xffffffff
}

//这种情况适合特殊的限定逻辑，但递归耗时还是有的
func minReorder_v2(n int, connections [][]int) int {
	var dfs func(curConnList [][]int, reOrderNum int, hasTraverse map[int]bool) int
	dfs = func(curConnList [][]int, reOrderNum int, hasTraverse map[int]bool) int {
		if len(curConnList) == 0 {
			if len(hasTraverse) == n {
				return reOrderNum
			} else {
				return 0xffffffff
			}
		}

		nxtConnList := make([][]int, 0)
		for _, curConn := range curConnList {
			if hasTraverse[curConn[1]] {
				hasTraverse[curConn[0]] = true
			} else if hasTraverse[curConn[0]] {
				hasTraverse[curConn[1]] = true
				reOrderNum++
			} else {
				nxtConnList = append(nxtConnList, curConn)
			}
		}

		return dfs(nxtConnList, reOrderNum, hasTraverse)
	}

	hasTraverse := make(map[int]bool)
	hasTraverse[0] = true
	return dfs(connections, 0, hasTraverse)
}

/*这种解法肯定是有必要的，能严格遍历出最小调整的边数*/
func minReorder_v1(n int, connections [][]int) int {
	var dfs func(curCity int, curConnList [][]int, reOrderNum int, hasTraverse map[int]bool) int
	dfs = func(curCity int, curConnList [][]int, reOrderNum int, hasTraverse map[int]bool) int {
		hasExist := hasTraverse[curCity]
		if !hasExist {
			hasTraverse[curCity] = true
		}
		defer func() {
			if !hasExist {
				delete(hasTraverse, curCity)
			}
		}()

		if len(curConnList) == 0 {
			if len(hasTraverse) == n {
				return reOrderNum
			} else {
				return 0xffffffff
			}
		}

		minRecorder := 0xffffffff
		for i, curConn := range curConnList {
			if hasTraverse[curConn[1]] {
				/*
					nxtConnList := make([][]int, len(curConnList)-1)
					for j := range nxtConnList {
						if j < i {
							nxtConnList[j] = curConnList[j]
						} else {
							nxtConnList[j] = curConnList[j+1]
						}
					}
				*/

				nxtConnList := make([][]int, 0)
				nxtConnList = append(curConnList[0:i], curConnList[i+1:]...)
				dfsOrderNum := dfs(curConn[0], nxtConnList, reOrderNum, hasTraverse)
				if dfsOrderNum < minRecorder {
					minRecorder = dfsOrderNum
				}
			}
		}

		for i, curConn := range curConnList {
			if hasTraverse[curConn[0]] {
				/*
					nxtConnList := make([][]int, len(curConnList)-1)
					for j := range nxtConnList {
						if j < i {
							nxtConnList[j] = curConnList[j]
						} else {
							nxtConnList[j] = curConnList[j+1]
						}
					}
				*/

				nxtConnList := make([][]int, 0)
				nxtConnList = append(curConnList[0:i], curConnList[i+1:]...)
				dfsOrderNum := dfs(curConn[1], nxtConnList, reOrderNum+1, hasTraverse)
				if dfsOrderNum < minRecorder {
					minRecorder = dfsOrderNum
				}
			}
		}

		return minRecorder
	}

	hasTraverse := make(map[int]bool)
	return dfs(0, connections, 0, hasTraverse)
}

func main() {
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	fmt.Printf("connections:%+v minReorder:%d\r\n", connections, minReorder(6, connections))

	connections = [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}
	fmt.Printf("connections:%+v minReorder:%d\r\n", connections, minReorder(5, connections))

	connections = [][]int{{1, 0}, {2, 0}}
	fmt.Printf("connections:%+v minReorder:%d\r\n", connections, minReorder(3, connections))
}
