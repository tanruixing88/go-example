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

func main() {
	bubbleSort([]int{4,5,3})
	bucketSort([]int{3,5,2,4,1}, 3)
	heapSort([]int{3,5,2,4,1})
}
