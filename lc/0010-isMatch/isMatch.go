package main

import "fmt"

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	match := func(i int, j int) bool {
		if i == 0 {
			return false
		}

		return s[i-1] == p[j-1] || p[j-1] == '.'
	}

	f := make([][]bool, len(s)+1)
	for i := range f {
		f[i] = make([]bool, len(p)+1)
	}

	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				//判断 s 并不匹配 c*s
				f[i][j] = f[i][j-2]

				//判断正则匹配了 s s*
				if match(i,j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else {
				if match(i, j) {
					f[i][j] = f[i-1][j-1]
				}
			}
		}
	}

	return f[m][n]
}

func main() {
	s := ""
	p := ""
	fmt.Printf("s:%s p:%s match:%t\r\n", s, p, isMatch(s, p))

	s = "abb"
	p = "c*a*b"
	fmt.Printf("s:%s p:%s match:%t\r\n", s, p, isMatch(s, p))

	s = "mississippi"
	p = "mis*is*p*."
	fmt.Printf("s:%s p:%s match:%t\r\n", s, p, isMatch(s, p))

	s = "dfadab"
	p = ".*ab"
	fmt.Printf("s:%s p:%s match:%t\r\n", s, p, isMatch(s, p))
}
