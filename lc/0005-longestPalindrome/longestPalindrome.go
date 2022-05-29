package main

import "fmt"

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		//奇数
		l := i
		h := i
		for l >= 0 && h < len(s) {
			if s[l] != s[h] {
				break
			}

			if h - l + 1 > len(longest) {
				longest = string([]byte(s)[l:h+1])
			}
			l--
			h++
		}

		//偶数
		l = i
		h = i + 1
		for l >= 0 && h < len(s) {
			if s[l] != s[h] {
				break
			}

			if h - l + 1 > len(longest) {
				longest = string([]byte(s)[l:h+1])
			}

			l--
			h++
		}
	}

	return longest
}

func main() {
	s := "babad"
	l := longestPalindrome(s)
	fmt.Printf("s:%s l:%s \r\n", s, l)

	s = "cbbd"
	l = longestPalindrome(s)
	fmt.Printf("s:%s l:%s \r\n", s, l)
}
