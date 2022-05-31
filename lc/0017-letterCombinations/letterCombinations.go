package main

import "fmt"

func letterCombinations(digits string) []string {
	digitsLetterMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	preCombinations := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		letters, ok := digitsLetterMap[digits[i]]
		if !ok {
			return make([]string, 0)
		}

		curCombinations := make([]string, 0)
		if len(preCombinations) == 0 {
			for j := 0; j < len(letters); j++ {
				curCombinations = append(curCombinations, letters[j])
			}
		} else {
			for k := 0; k < len(preCombinations); k++ {
				for j := 0; j < len(letters); j++ {
					curCombinations = append(curCombinations, preCombinations[k] + letters[j])
				}
			}
		}
		preCombinations = curCombinations
	}

	return preCombinations
}

func main() {
	digits := "23"
	letterCombinationList := letterCombinations(digits)
	fmt.Printf("digit:%s letterCombinationList:%+v\r\n", digits, letterCombinationList)

	digits = ""
	letterCombinationList = letterCombinations(digits)
	fmt.Printf("digit:%s letterCombinationList:%+v\r\n", digits, letterCombinationList)

	digits = "2"
	letterCombinationList = letterCombinations(digits)
	fmt.Printf("digit:%s letterCombinationList:%+v\r\n", digits, letterCombinationList)

	digits = "234"
	letterCombinationList = letterCombinations(digits)
	fmt.Printf("digit:%s letterCombinationList:%+v\r\n", digits, letterCombinationList)

}
