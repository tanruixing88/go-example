package main

import "fmt"

func rotate(matrix [][]int)  {
	n := len(matrix)
	for cir := 0; cir < n / 2; cir++ {
		tmp := make([]int, n-cir)
		for i := cir; i < n-cir; i++ {
			tmp[i] = matrix[cir][i]
		}

		//left -> up
		for i := cir; i < n-cir; i++ {
			matrix[cir][n-1-i] = matrix[i][cir]
		}

		//down -> left
		for i := cir; i < n-cir; i++ {
			matrix[i][cir] = matrix[n-1-cir][i]
		}

		//right -> down
		for i := cir; i < n-cir; i++ {
			matrix[n-1-cir][i] = matrix[n-1-i][n-1-cir]
		}

		//tmp -> right
		for i := cir; i < n-cir; i++ {
			matrix[n-1-i][n-1-cir] = tmp[n-1-i]
		}
	}
}

func main() {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Printf("matrix:%+v", matrix)
	rotate(matrix)
	fmt.Printf("after rotate:%+v\r\n", matrix)

	matrix = [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	fmt.Printf("matrix:%+v", matrix)
	rotate(matrix)
	fmt.Printf("after rotate:%+v\r\n", matrix)
}
