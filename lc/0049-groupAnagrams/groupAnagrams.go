package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		//给str进行字符串排序
		key := []byte(strs[i])
		for j := 0; j < len(key); j++{
			for k := 0; k < len(key) - 1 - j; k++ {
				if key[k+1] < key[k] {
					key[k+1], key[k] = key[k], key[k+1]
				}
			}
		}

		_, ok := groupMap[string(key)]
		if ok {
			groupMap[string(key)] = append(groupMap[string(key)], strs[i])
		} else {
			groupMap[string(key)] = []string{strs[i]}
		}
	}

	groupList := make([][]string, 0)
	for _, group := range groupMap {
		groupList = append(groupList, group)
	}

	return groupList
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupList := groupAnagrams(strs)
	fmt.Printf("strs:%+v groupList:%+v\r\n", strs, groupList)

	strs = []string{""}
	groupList = groupAnagrams(strs)
	fmt.Printf("strs:%+v groupList:%+v\r\n", strs, groupList)

	strs = []string{"a"}
	groupList = groupAnagrams(strs)
	fmt.Printf("strs:%+v groupList:%+v\r\n", strs, groupList)
}
