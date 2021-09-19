package main

import "fmt"

func binarySearch(a []int, e int) int {
	if len(a) == 0 {
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

//a first half is asc, second half is asc, first half value gt second half
func binarySearchAscAsc(a []int, e int) int {
	if len(a) == 0 {
		return -1
	}

	l := 0
	h := len(a) - 1

	for l <= h {
		m := (l + h) / 2
		//first should judge e op a[0]
		if e > a[0] {
			if a[m] < a[0] {
				h = m -1
			} else if a[m] == e {
				return m
			} else {
				if e > a[m] {
					l = m + 1
				} else {
					h = m - 1
				}
			}
		} else if e == a[0] {
			return 0
		} else if e == a[len(a)-1] {
			return len(a) - 1
		} else {
			if a[m] < a[len(a)-1] {
				if a[m] < e {
					l = m + 1
				} else {
					h = m - 1
				}
			} else if a[m] == e {
				return m
			} else {
				l = m + 1
			}
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

	a = []int{1,2,3,4,5,-6,-7,-8}
	i = binarySearchAscAsc(a, -7)
	fmt.Printf("index:%d \r\n", i)
}

