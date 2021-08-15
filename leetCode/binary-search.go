package main

import "fmt"

func binarySearch(a []int, e int) int {
	if len(a) < 1 {
		return -1
	}

	l := 0
	h := len(a) - 1

	for l <= h {
		m := (l+h) / 2

		if a[m] < e {
			l = m + 1
		} else if a[m] > e {
			h = m - 1
		} else {
			return m
		}
	}

	return -1
}

func main()  {
	a := []int{}
	i := binarySearch(a, 1)
	fmt.Printf("index:%d \r\n", i)

	a = []int{1,3}
	i = binarySearch(a, 2)
	fmt.Printf("index:%d \r\n", i)

	a = []int{1,2,3,5,6,7,8}
	i = binarySearch(a, 4)
	fmt.Printf("index:%d \r\n", i)

	a = []int{1,2,3,4,5,6,7,8}
	i = binarySearch(a, 7)
	fmt.Printf("index:%d \r\n", i)
}

