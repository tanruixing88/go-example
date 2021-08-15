package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	p := a[0]
	l := 1
	h := len(a) - 1
	for l <= h {
		if a[l]	<= p {
			l++
			continue
		}

		if a[h] >= p {
			h--
			continue
		}

		tmp := a[l]
		a[l] = a[h]
		a[h] = tmp
	}

	if l > 1 {
		quickSort(a[1:l])
	}

	if h < len(a) -1 {
		quickSort(a[h+1:])
	}

	if !(l == 1 && a[l] > p) {
		for i := 1; i < l; i++ {
			a[i-1] = a[i]
		}
		a[l-1] = p
	}

	return a
}

func main() {
	var a []int
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{1}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{1,2}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{2,1}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{1,2,3}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{3,2,1}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{3,2,1,4}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{5,0,3,2,1,4,7,6}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)

	a = []int{4,1,9,0,3,5,2,10,6,7}
	quickSort(a)
	fmt.Printf("%+v\r\n", a)
}
