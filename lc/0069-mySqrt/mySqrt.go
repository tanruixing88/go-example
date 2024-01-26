package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	l := 1
	h := x

	for l < h {
		m := (l + h) / 2
		if m*m <= x && (m+1)*(m+1) >= x {
			if (m+1)*(m+1) == x {
				return m + 1
			}

			return m
		} else if (m+1)*(m+1) > x {
			h = m
		} else {
			l = m
		}
	}

	return 1
}

func main() {
	/*
		x := 2
		fmt.Printf("x:%d sqrt:%d \r\n", x, mySqrt(x))
		x = 4
		fmt.Printf("x:%d sqrt:%d \r\n", x, mySqrt(x))
		x = 8
		fmt.Printf("x:%d sqrt:%d \r\n", x, mySqrt(x))
	*/
	x := 401
	fmt.Printf("x:%d sqrt:%d \r\n", x, mySqrt(x))
}
