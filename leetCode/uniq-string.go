package main

import "fmt"

//字符串字符均不同
//https://github.com/lifei6671/interview-go/blob/master/question/q002.md

func uniqString(s1, s2 string) bool {
	bitMap := make([]int64, 2)
	for i := 0; i < len(s1); i++ {
		v := uint64(s1[i])
		if v >= 64 {
			bitMap[1] |= int64(1 << (v - 64))
		} else {
			bitMap[0] |= int64(1 << v)
		}
	}

	for i := 0; i < len(s2); i++ {
		v := uint64(s2[i])
		if v >= 64 {
			if (int64(1 << (v - 64)) & bitMap[1]) > 0 {
				return false
			}
		} else {
			if (int64(1 << v) & bitMap[0]) > 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	r := uniqString("", "")
	fmt.Printf("ret:%t \r\n", r)

	r = uniqString("abc", "123")
	fmt.Printf("ret:%t \r\n", r)

	r = uniqString("def", "fig")
	fmt.Printf("ret:%t \r\n", r)
}
