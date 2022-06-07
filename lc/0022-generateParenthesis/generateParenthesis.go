package main

import "fmt"

func generateParenthesis(n int) []string {
	curExistMap := make(map[string]bool)
	curExistMap["()"] = true
	for i := 1; i < n; i++ {
		nxtExistMap := make(map[string]bool)
		for curKey := range curExistMap {
			for j := 0; j < len(curKey); j++ {
				for k := j; k < len(curKey); k++ {
					nxtKey := curKey[:j] + "(" + curKey[j:k] + ")" + curKey[k:]
					_, ok := nxtExistMap[nxtKey]
					if !ok {
						nxtExistMap[nxtKey] = true
					}
				}
			}
		}
		curExistMap = nxtExistMap
	}

	parenthesisList := make([]string, 0)
	for key := range curExistMap {
		parenthesisList = append(parenthesisList, key)
	}

	return parenthesisList
}

func main() {
	n := 1
	parenthesisList := generateParenthesis(n)
	fmt.Printf("n:%d parenthesisList:%+v\r\n", n, parenthesisList)

	n = 2
	parenthesisList = generateParenthesis(n)
	fmt.Printf("n:%d parenthesisList:%+v\r\n", n, parenthesisList)

	n = 3
	parenthesisList = generateParenthesis(n)
	fmt.Printf("n:%d parenthesisList:%+v\r\n", n, parenthesisList)
}
