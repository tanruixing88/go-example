package main

import "fmt"

//声明结构对于一堆（（（（（可以有效解决处理长度

type leftFrequency struct {
	LeftByte byte
	LeftCount int
}

func isValid(s string) bool {
	stack := make([]leftFrequency, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			if len(stack) > 0 && stack[len(stack)-1].LeftByte == s[i] {
				stack[len(stack)-1].LeftCount++
			} else if len(stack) == 0 || stack[len(stack)-1].LeftByte != s[i] {
				stack = append(stack, leftFrequency{s[i], 1})
			}
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			var matchByte byte
			if s[i] == ')' {
				matchByte = '('
			} else if s[i] == ']' {
				matchByte = '['
			} else if s[i] == '}' {
				matchByte = '{'
			}
			if len(stack) == 0 || stack[len(stack)-1].LeftByte != matchByte {
				return false
			} else {
				if stack[len(stack)-1].LeftCount == 1 {
					stack = stack[:len(stack)-1]
				} else {
					stack[len(stack)-1].LeftCount--
				}
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "()"
	r := isValid(s)
	fmt.Printf("s:%s r:%t\r\n", s, r)

	s = "()[]{}"
	r = isValid(s)
	fmt.Printf("s:%s r:%t\r\n", s, r)

	s = "(]"
	r = isValid(s)
	fmt.Printf("s:%s r:%t\r\n", s, r)

	s = "([)]"
	r = isValid(s)
	fmt.Printf("s:%s r:%t\r\n", s, r)

	s = "{[]}"
	r = isValid(s)
	fmt.Printf("s:%s r:%t\r\n", s, r)
}
