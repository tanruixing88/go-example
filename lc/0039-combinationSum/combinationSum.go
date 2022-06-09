package main

import (
	"fmt"
	"sort"
)

//层序遍历方式需要更多的内存来保留中间状态，还需要去除混杂路径的组合重复问题
func combinationSum1(candidates []int, target int) [][]int {
	combinations := make([][]int, 0)
	uniqMap := make(map[string]bool)
	curTargetPrefixMap := make(map[int][][]int)
	curTargetList := []int{target}
	for len(curTargetList) > 0 {
		nxtTargetPrefixMap := make(map[int][][]int)
		for l := 0; l < len(curTargetList); l++ {
			for i := 0; i < len(candidates); i++ {
				if candidates[i] < curTargetList[l] {
					nxtTarget := curTargetList[l] - candidates[i]
					prefixList, ok := curTargetPrefixMap[curTargetList[l]]
					nxtPrefixList := make([][]int, 0)
					if ok {
						for j := 0; j < len(prefixList); j++ {
							tmpList := make([]int, 0)
							for k := 0; k < len(prefixList[j]); k++ {
								tmpList = append(tmpList, prefixList[j][k])
							}
							tmpList = append(tmpList, candidates[i])
							nxtPrefixList = append(nxtPrefixList, tmpList)
						}
					} else {
						nxtPrefixList = [][]int{{candidates[i]}}
					}

					_, ok = nxtTargetPrefixMap[nxtTarget]
					if ok {
						allNxtPrefixList := make([][]int, 0)
						allKeyExistMap := make(map[string]bool)
						for j := 0; j < len(nxtTargetPrefixMap[nxtTarget]); j++ {
							sort.Ints(nxtTargetPrefixMap[nxtTarget][j])
							allKey := ""
							for x := 0; x < len(nxtTargetPrefixMap[nxtTarget][j]); x++ {
								allKey += fmt.Sprintf("%d", nxtTargetPrefixMap[nxtTarget][j])
							}
							_, allOk :=  allKeyExistMap[allKey]
							if allOk {
								continue
							} else {
								allKeyExistMap[allKey] = true
								allNxtPrefixList = append(allNxtPrefixList, nxtTargetPrefixMap[nxtTarget][j])
							}
						}

						for j := 0; j < len(nxtPrefixList); j++ {
							sort.Ints(nxtPrefixList[j])
							allKey := ""
							for x := 0; x < len(nxtPrefixList[j]); x++ {
								allKey += fmt.Sprintf("%d", nxtPrefixList[j])
							}
							_, allOk := allKeyExistMap[allKey]
							if allOk {
								continue
							} else {
								allKeyExistMap[allKey] = true
								allNxtPrefixList = append(allNxtPrefixList, nxtPrefixList[j])
							}
						}
						nxtTargetPrefixMap[nxtTarget] = allNxtPrefixList
					} else {
						nxtTargetPrefixMap[nxtTarget] = nxtPrefixList
					}

					/*
					fmt.Printf("\r\ncurTarget:%d candidate:%d curTargetList:%+v nxtTargetPrefixMap:%+v \r\n",
						curTargetList[l], candidates[i], curTargetList, nxtTargetPrefixMap)
					 */

				} else if candidates[i] == curTargetList[l] {
					prefixList, ok := curTargetPrefixMap[curTargetList[l]]
					if ok {
						for j := 0; j < len(prefixList); j++ {
							combination := make([]int, 0)
							for _, v := range prefixList[j] {
								combination = append(combination, v)
							}
							combination = append(combination, candidates[i])
							sort.Ints(combination)
							key := ""
							sum := 0
							for k := 0; k < len(combination); k++ {
								key += fmt.Sprintf("%d", combination[k])
								sum += combination[k]
							}
							if sum != target {
								continue
							}
							if uniqMap[key] {
								continue
							} else {
								uniqMap[key] = true
							}
							combinations = append(combinations, combination)
						}
					} else {
						if candidates[i] == target {
							combinations = append(combinations, []int{candidates[i]})
						}
					}
				} else {
					//忽略
				}
			}
		}
		curTargetList = make([]int, 0)
		for nxtTarget := range nxtTargetPrefixMap {
			curTargetList = append(curTargetList, nxtTarget)
		}
		curTargetPrefixMap = nxtTargetPrefixMap
		//fmt.Printf("\r\ncurTargetPrefixMap:%+v\r\n", curTargetPrefixMap)
	}

	return combinations
}


//见官方题解，https://leetcode.cn/problems/combination-sum/solution/zu-he-zong-he-by-leetcode-solution/
//不需要额外的遍历，仅仅需要一个全局的comb栈队列，走一遍遍历即可
func combinationSum(candidates []int, target int) [][]int {
	comb := []int{}
	ans := make([][]int, 0)
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}


func main() {
	/*
	candidates := []int{2,3,6,7}
	target := 7
	combinations := combinationSum(candidates, target)
	fmt.Printf("candidates:%+v target:%d combinations:%+v\r\n", candidates, target, combinations)

	candidates = []int{2,3,5}
	target = 8
	combinations = combinationSum(candidates, target)
	fmt.Printf("candidates:%+v target:%d combinations:%+v\r\n", candidates, target, combinations)

	candidates = []int{2}
	target = 1
	combinations = combinationSum(candidates, target)
	fmt.Printf("candidates:%+v target:%d combinations:%+v\r\n", candidates, target, combinations)

	candidates = []int{2,7,6,3,5,1}
	target = 9
	combinations = combinationSum(candidates, target)
	fmt.Printf("candidates:%+v target:%d combinations:%+v\r\n", candidates, target, combinations)

	 */

	candidates := []int{7,3,2}
	target := 18
	combinations := combinationSum(candidates, target)
	fmt.Printf("candidates:%+v target:%d combinations:%+v\r\n", candidates, target, combinations)
}
