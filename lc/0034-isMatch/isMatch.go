package main

import "fmt"

func isMatch(s string, p string) bool {
	a := make([][]bool, len(s)+1)
	for i := range a {
		a[i] = make([]bool, len(p)+1)
	}

	a[0][0] = true

	if len(p) > 0 && p[0] == '*' {
		a[0][1] = true
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				a[0][j] = a[0][j] || a[0][j-1]
			}
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] != '*' {
				if s[i-1] == p[j-1] || p[j-1] == '?' {
					a[i][j] = a[i-1][j-1]
				}
			} else {
				a[i][j] = a[i][j-1] || a[i-1][j] || a[i-1][j-1]
			}
		}
	}

	fmt.Printf("\t")
	fmt.Printf("\t")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\t", s[i])
	}
	fmt.Printf("\r\n")
	for j := 0; j <= len(p); j++ {
		if j > 0 {
			fmt.Printf("%c\t", p[j-1])
		} else {
			fmt.Printf("\t")
		}
		for i := 0; i <= len(s); i++ {
			fmt.Printf("%t\t", a[i][j])
		}
		fmt.Printf("\r\n")
	}

	return a[len(s)][len(p)]
}

func main() {
	s := "aa"
	p := "a"
	fmt.Printf("s:%s p:%s isMatch:%t\r\n", s, p, isMatch(s, p))

	s = "aa"
	p = "*"
	fmt.Printf("s:%s p:%s isMatch:%t\r\n", s, p, isMatch(s, p))

	s = "cb"
	p = "?a"
	fmt.Printf("s:%s p:%s isMatch:%t\r\n", s, p, isMatch(s, p))

	s = "adceb"
	p = "*a*b"
	fmt.Printf("s:%s p:%s isMatch:%t\r\n", s, p, isMatch(s, p))

	s = "miss"
	p = "m??*s"
	fmt.Printf("s:%s p:%s isMatch:%t\r\n", s, p, isMatch(s, p))

	s = ""
	p = "******"
	fmt.Printf("s:%s p:%s isMatch:%t\r\n", s, p, isMatch(s, p))
}
