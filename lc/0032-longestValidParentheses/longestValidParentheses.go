package main

import "fmt"

func longestValidParentheses(s string) int {
	maxLen := 0
	leftMatchIdx := 0
	matchNumList := make([]int, len(s))
	matchIdxStack := make([]int, 0)
	for i, b := range s {
		if b == '(' {
			matchIdxStack = append(matchIdxStack, i)
		} else if b == ')' {
			if len(matchIdxStack) > 0 {
				leftMatchIdx = matchIdxStack[len(matchIdxStack)-1]
				matchNumList[i] = i - leftMatchIdx + 1
				if leftMatchIdx > 0 {
					matchNumList[i] += matchNumList[leftMatchIdx-1]
				}
				matchIdxStack = matchIdxStack[:len(matchIdxStack)-1]

				if matchNumList[i] > maxLen {
					maxLen = matchNumList[i]
				}
			}
		}
	}

	//fmt.Printf("s:%s matchNumList:%+v\r\n", s, matchNumList)
	return maxLen
}

func longestValidParentheses1(s string) int {
	maxAns := 0
	//最后一个没有被匹配的右括号的下标
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}

	return maxAns
}

func longestValidParentheses3(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "(()"
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = "))("
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = ""
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = ")()()("
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = "()(()"
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = "(()((()))"
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = "(()())"
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = ")()())()()("
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses(s))
	s = "()"
	fmt.Printf("s:%s longestValidParentheses:%d\r\n", s, longestValidParentheses1(s))
}
