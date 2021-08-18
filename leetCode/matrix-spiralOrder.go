package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	el := []int{}

	if matrix == nil || len(matrix) < 1 || len(matrix[0]) < 1 {
		return el
	}

	initLen := len(matrix[0])
	for _, matrixD1 := range matrix {
		if initLen != len(matrixD1) {
			fmt.Printf("error param \r\n")
			return el
		}
	}

	si := 0
	ei := len(matrix) - 1
	sj := 0
	ej := len(matrix[0]) - 1
	for true {
		for j := sj; j <= ej; j++ {
			el = append(el, matrix[si][j])
		}
		if si == ei {
			break
		} else {
			si++
		}

		for i := si; i <= ei; i++ {
			el = append(el, matrix[i][ej])
		}
		if sj == ej {
			break
		} else {
			ej--
		}

		for j := ej; j >= sj; j-- {
			el = append(el, matrix[ei][j])
		}
		if si == ei {
			break
		} else {
			ei--
		}

		for i := ei; i >= si; i-- {
			el = append(el, matrix[i][sj])
		}
		if sj == ej {
			break
		} else {
			sj++
		}
	}

	return el
}

func main() {
	matrix := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	el := spiralOrder(matrix) //except [1,2,3,6,9,8,7,4,5]
	fmt.Printf("%+v\r\n", el)

	matrix = [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}}
	el = spiralOrder(matrix) //except [1,2,3,4,8,12,11,10,9,5,6,7]
	fmt.Printf("%+v\r\n", el)
}
