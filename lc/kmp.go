package main

import "fmt"

//详细解释：https://dal-s.github.io/2019/07/09/golang%E5%AE%9E%E7%8E%B0KMP/

func kmp(s string, p string) int {
	m := len(s)
	n := len(p)
	next := make([]int, n)
	next[0] = -1
	next[1] = 0

	/* 这种方案需要区分奇偶数*/
	/*
	for j := 2; j < n; j++ {
		lj := 0
		hj := 0

		if j % 2 == 0 {
			lj = (j - 1) / 2
			hj = (j - 1) / 2 + 1
		} else {
			lj = (j - 1) / 2 - 1
			hj = (j - 1) / 2 + 1
		}

		for lj >= 0 && hj < n {
			tlj := 0
			thj := hj
			for tlj <= lj {
				if p[tlj] == p[thj] {
					tlj++
					thj++
				} else {
					break
				}
			}

			if tlj > lj {
				next[j] = lj + 1
				break
			} else {
				lj--
				hj++
			}
		}
	}
	*/

	//直接用子字符串切割，子字符串比较
	for j := 2; j < n; j++ {
		for k := j / 2; k >= 0; k-- {
			if p[0:k] == p[j-k:j] {
				next[j]	= k
				break
			}
		}
	}
	fmt.Printf("next:%+v\r\n", next)
	//至此已经构建完毕next数组

	j := 0
	i := 0
	for i < m && j < n {
		if j == -1 || s[i] == p[j] {
			i++
			j++
			continue
		} else {
			j = next[j]
		}
	}

	if j == n {
		return i - j
	} else {
		return -1
	}
}

func main() {
	s := "addafdasjklf"
	p := "dafdas"
	i := kmp(s, p)
	fmt.Printf("pos:%d\r\n", i)
}
