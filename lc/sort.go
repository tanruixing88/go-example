package main

import "fmt"

//冒泡排序
//https://www.runoob.com/w3cnote/bubble-sort.html
func bubbleSort1(a []int) {
	for i := len(a)-1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("bubbleSort:%+v\r\n", a)
}

func bubbleSort(a []int) {
	for i := 0; i < len(a) - 1; i++ {
		for j := 0; j < len(a) - 1 - i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("bubbleSort:%+v\r\n", a)
}

//桶排序
//https://www.runoob.com/w3cnote/bucket-sort.html
func bucketSort(a []int, bucketSize int) {
	if len(a) <= 1 || bucketSize < 1 {
		return
	}

	n := len(a)
	minV := a[0]
	maxV := minV
	for i := 0; i < n; i++ {
		if a[i] > maxV {
			maxV = a[i]
		}

		if a[i] < minV {
			minV = a[i]
		}
	}

	bucketCnt := (maxV - minV + 1) / bucketSize + 1

	allBuckets := make([][]int, bucketCnt)
	for i := 0; i < len(allBuckets); i++ {
		allBuckets[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		idx := (a[i] - minV + 1) / bucketSize
		allBuckets[idx] = append(allBuckets[idx], a[i])
	}

	for i := 0; i < len(allBuckets); i++ {
		for j := 0; j < len(allBuckets[i]) - 1; j++ {
			for k := 0; k < len(allBuckets[i]) - 1 - j; k++ {
				if allBuckets[i][k]	> allBuckets[i][k+1] {
					allBuckets[i][k], allBuckets[i][k+1] = allBuckets[i][k+1], allBuckets[i][k]
				}
			}
		}
	}

	k := 0
	for i := 0; i < len(allBuckets); i++ {
		for j := 0; j < len(allBuckets[i]); j++ {
			a[k] = allBuckets[i][j]
			k++
		}
	}

	fmt.Printf("bucketSort:%+v\r\n", a)
}

//堆排序
//https://www.runoob.com/w3cnote/heap-sort.html
//为何为j-1/2？初始index为0 1 2 。。。 其中0为堆顶，1,2为堆下一层，所以用 j-1/2使得1和2下标均能对应0下标
func heapSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n - 1; i++ {
		for j := n - 1; j > i; j-- {
			if a[j] < a[(j-1)/2] {
				a[j], a[(j-1)/2] = a[(j-1)/2], a[j]
			}
		}
	}

	fmt.Printf("heapSort:%+v\r\n", a)
}

func quickSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	if n == 2 {
		if a[1] > a[0] {
			a[0], a[1] = a[1], a[0]
		}
		return
	}

	p := a[0]
	l := 1
	h := n-1

	for l < h {
		if a[l] <= p {
			l++
			continue
		}

		if a[h] >= p {
			h--
			continue
		}

		a[l], a[h] = a[h], a[l]
	}

	if a[l] > p {
		l--
	} else if a[h] < p {
		h++
	} else {
		l--
	}

	if l > 1 {
		quickSort(a[1:h])
	}

	if h < n - 1 {
		quickSort(a[h:])
	}

	if a[1] < p {
		for i := 1; i <= l; i++ {
			a[i-1] = a[i]
		}
		a[l] = p
	}

	fmt.Printf("a:%+v\r\n", a)
}

func main() {
	bubbleSort([]int{4,5,3})
	bucketSort([]int{3,5,2,4,1}, 3)
	heapSort([]int{3,5,2,4,1})
	quickSort([]int{3,5,2,4,1})
	quickSort([]int{3,5})
	quickSort([]int{11,1,9,0,3,5,2,10,6,7})
}
