package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	existMap := make(map[byte]int)
	maxLen := 0
	for i := 0; i < len(s); i++ {
		idx, ok := existMap[s[i]]
		if ok {
			for k, v := range existMap {
				if v < idx {
					delete(existMap, k)
				}
			}
		}

		existMap[s[i]] = i

		if maxLen < len(existMap) {
			maxLen = len(existMap)
		}
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	r := lengthOfLongestSubstring(s)
	fmt.Printf("r:%d\r\n", r)

	s = "bbbbb"
	r = lengthOfLongestSubstring(s)
	fmt.Printf("r:%d\r\n", r)

	s = "pwwkew"
	r = lengthOfLongestSubstring(s)
	fmt.Printf("r:%d\r\n", r)

	s = "a"
	r = lengthOfLongestSubstring(s)
	fmt.Printf("r:%d\r\n", r)
}
