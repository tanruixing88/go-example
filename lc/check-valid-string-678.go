package main

import "fmt"

// https://leetcode-cn.com/problems/valid-parenthesis-string/
func checkValidString(s string) bool {
	if len(s) == 0 {
		return true
	}

	l := 0
	h := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
			h++
		} else if s[i] == ')' {
			if l > 0 {
				l--
			}
			h--
			if h < 0 {
				return false
			}
		} else if s[i] == '*' {
			if l > 0 {
				l--
			}
			h++
		}
	}

	return l == 0
}
// 要用map格式，不要用list
func checkValidString1(s string) bool {
	if len(s) == 0 {
		return true
	}

	leftMap := make(map[int]bool)
	leftMap[0] = true
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			tmpMap := make(map[int]bool)
			for j := range leftMap {
				tmpMap[j+1] = true
			}
			leftMap = tmpMap
		} else if s[i] == ')' {
			tmpMap := make(map[int]bool)
			for j := range leftMap {
				if j - 1 >= 0 {
					tmpMap[j-1] = true
				}
			}
			if len(tmpMap) == 0 {
				return false
			}
			leftMap = tmpMap
		} else if s[i] == '*' {
			tmpMap := make(map[int]bool)
			for j := range leftMap {
				tmpMap[j] = true
				tmpMap[j + 1] = true
				if j - 1 >= 0 {
					tmpMap[j - 1] = true
				}
			}
			leftMap = tmpMap
		} else {
			return false
		}
	}

	if leftMap[0] {
		return true
	}

	return false
}

func main() {
	s := "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"
	r := checkValidString(s)
	fmt.Printf("r:%t\r\n", r)

}
