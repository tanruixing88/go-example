package main

import "fmt"
// https://leetcode-cn.com/problems/shortest-palindrome/solution/zui-duan-hui-wen-chuan-by-leetcode-solution/
//官方KMP
func shortestPalindrome(s string) string {
	n := len(s)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i - 1]
		for j != -1 && s[j + 1] != s[i] {
			j = fail[j]
		}
		if s[j + 1] == s[i] {
			fail[i] = j + 1
		}
	}
	fmt.Printf("s:%s  fail:%+v\r\n", s, fail)

	best := -1
	for i := n - 1; i >= 0; i-- {
		for best != -1 && s[best + 1] != s[i] {
			best = fail[best]
		}
		if s[best + 1] == s[i] {
			best++
		}
	}
	add := ""
	if best != n - 1 {
		add = s[best + 1:]
	}
	b := []byte(add)
	for i := 0; i < len(b) / 2; i++ {
		b[i], b[len(b) - 1 -i] = b[len(b) - 1 -i], b[i]
	}
	return string(b) + s
}


//普通方法
func shortestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}

	getPalindrome := func (l int, h int) string {
		tl := l
		th := h
		for tl >= 0 && th < len(s) {
			if s[tl] == s[th] {
				tl--
				th++
			} else {
				if l == h {
					l--
				} else {
					h--
				}

				tl = l
				th = h
			}
		}

		if tl < 0 {
			tl++
			th--
		}

		bytes := make([]byte, 0)
		for i := len(s) - 1; i > th; i-- {
			bytes = append(bytes, s[i])
		}

		return string(bytes) + s
	}

	mid := (len(s) - 1) / 2
	l := mid
	h := mid
	if len(s) % 2 == 0 {
		return getPalindrome(l, h+1)
	} else {
		return getPalindrome(l, h)
	}
}

func main() {
	s := "aaaab"
	r := shortestPalindrome(s)
	fmt.Printf("ret:%s\r\n", r)

	s = "abababc"
	r = shortestPalindrome(s)
	fmt.Printf("ret:%s\r\n", r)

	s = "abcabcd"
	r = shortestPalindrome(s)
	fmt.Printf("ret:%s\r\n", r)

	s = "aacecaaa"
	r = shortestPalindrome(s)
	fmt.Printf("ret:%s\r\n", r)

	s = "ba"
	r = shortestPalindrome(s)
	fmt.Printf("ret:%s\r\n", r)

	s = "abcdabd"
	r = shortestPalindrome(s)
	fmt.Printf("ret:%s\r\n", r)
}