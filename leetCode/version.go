package main

import "fmt"

func versionCmp(v1 string, v2 string) int {
	v1L := make([]int, 0)
	vNum := 0
	for i := 0; i < len(v1); i++ {
		if v1[i] >= '0' && v1[i] <= '9' {
			vNum = vNum * 10 + int(v1[i] - '0')
		} else if v1[i] == '.' {
			v1L = append(v1L, vNum)
			vNum = 0
		}
	}
	v1L = append(v1L, vNum)

	v2L := make([]int, 0)
	vNum = 0
	for i := 0; i < len(v2); i++ {
		if v2[i] >= '0' && v2[i] <= '9' {
			vNum = vNum * 10 + int(v2[i] - '0')
		} else if v2[i] == '.' {
			v2L = append(v2L, vNum)
			vNum = 0
		}
	}
	v2L = append(v2L, vNum)

	j := 0
	for j < len(v1L) && j < len(v2L) {
		if v1L[j] > v2L[j] {
			return 1
		} else if v1L[j] < v2L[j] {
			return -1
		}
		j++
	}

	if len(v1L)	> len(v2L) {
		return 1
	} else if len(v1L) < len(v2L) {
		return -1
	}

	return 0
}

func main() {
	v1 := "1.02.11"
	v2 := "2.14.4"
	fmt.Printf("ret:%d\r\n", versionCmp(v1, v2))
	v1 = "1.02.11"
	v2 = "2.14.4"
	fmt.Printf("ret:%d\r\n", versionCmp(v1, v2))
}
