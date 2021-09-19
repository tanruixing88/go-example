package main

import "fmt"

//接雨水
//https://leetcode-cn.com/problems/trapping-rain-water/
//自己方式实现
func trap1(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	topIdxList := make([]int, 0)
	lTopIdx := 0
	hTopIdx := 0

	ascFlag := true
	for i := 1; i < len(height); i++ {
		if height[i] > height[i-1] {
			lTopIdx = i
			hTopIdx = i
			ascFlag = true
			//fmt.Printf("lTopIdx:%d hTopIdx:%d\r\n", lTopIdx, hTopIdx)
		} else if height[i] < height[i-1] {
			if !ascFlag {
				continue
			}
			//fmt.Printf("i:%d height:%d\r\n", i, height[i])
			for j := lTopIdx; j <= hTopIdx; j++ {
				topIdxList = append(topIdxList, j)
			}
			ascFlag = false
		} else {
			hTopIdx++
		}
	}

	fmt.Printf("topIdxList:%v\r\n", topIdxList)
	for j := lTopIdx; j <= hTopIdx; j++ {
		topIdxList = append(topIdxList, j)
	}

	if len(topIdxList) <= 1 {
		return 0
	}

	trapNum := 0
	curTopIdx := topIdxList[0]
	maxTopIdx := -1
	i := 0
	for i < len(topIdxList) {
		addFlag := false
		maxK := -1
		for k := i+1; k < len(topIdxList); k++ {
			if height[topIdxList[k]] >= height[curTopIdx] {
				for j := curTopIdx; j < topIdxList[k]; j++ {
					if height[curTopIdx] > height[j] {
						trapNum += height[curTopIdx] - height[j]
					}
				}
				fmt.Printf("curTopIdx:%d topk:%d d:%d\r\n", curTopIdx, topIdxList[k], trapNum)
				curTopIdx = topIdxList[k]
				addFlag = true
				break
			} else {
				if maxTopIdx == -1 || height[topIdxList[k]] > height[maxTopIdx] {
					maxTopIdx = topIdxList[k]
					maxK = k
				}
			}
		}

		if addFlag {
			maxTopIdx = -1
			i++
		} else {
			for j := curTopIdx+1; j < maxTopIdx; j++ {
				if height[j] > height[maxTopIdx] {
					continue
				}
				trapNum += height[maxTopIdx] - height[j]
			}
			if maxTopIdx == -1 {
				break
			} else {
				curTopIdx = maxTopIdx
			}
			i = maxK
			maxTopIdx = -1
		}
	}

	return trapNum
}

// 类似单调栈的解决方法
func trap2(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	trapNum := 0
	l := 0
	lMaxV := height[l]
	pLMaxV := height[l]
	h := len(height) - 1
	hMaxV := height[h]
	pHMaxV := hMaxV
	lastLevel := 0
	level := 0

	for l < h {
		fmt.Printf("l:%d h:%d\r\n", l, h)
		if lMaxV <= height[l+1] {
			l++
			lMaxV = height[l]
			continue
		} else {
			if pLMaxV > 0 && pLMaxV == lastLevel && height[l] <= pLMaxV {
				l++
				continue
			}
		}

		if hMaxV <= height[h-1] {
			h--
			hMaxV = height[h]
			continue
		} else {
			if pHMaxV > 0 && pHMaxV == lastLevel && height[h] <= pHMaxV {
				h--
				continue
			}
		}

		pLMaxV = lMaxV
		pHMaxV = hMaxV

		if lMaxV < hMaxV {
			level = lMaxV
		} else {
			level = hMaxV
		}

		for i := l; i <= h; i++ {
			if height[i] < level {
				if height[i] > lastLevel {
					trapNum += level - height[i]
				} else {
					trapNum += level - lastLevel
				}
			}
		}

		fmt.Printf("l:%d h:%d level:%d lastLevel:%d trapNum:%d\r\n", l, h, level, lastLevel, trapNum)
		lastLevel = level
	}

	return trapNum
}

//类似双指针
func trap(height []int) int {
	l := 0
	h := len(height) - 1
	lMaxV := height[l]
	hMaxV := height[h]
	trapNum := 0
	for l < h {
		if lMaxV <= hMaxV {
			if height[l+1] > lMaxV {
				lMaxV = height[l+1]
			} else {
				trapNum += lMaxV - height[l+1]
			}
			l++
		} else {
			if height[h-1] > hMaxV {
				hMaxV = height[h-1]
			} else {
				trapNum += hMaxV - height[h-1]
			}
			h--
		}
	}

	return trapNum
}


func main() {
	r := trap([]int{0,1,0,2})
	fmt.Printf("r:%d\r\n", r)
	r = trap([]int{8,5,4,1,8,9,3,0,0})
	fmt.Printf("r:%d\r\n", r)
	r = trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})
	fmt.Printf("r:%d\r\n", r)
	/*
	r = trap([]int{4,2,0,3,2,5})
	fmt.Printf("r:%d\r\n", r)
	r = trap([]int{5,4,1,2})
	fmt.Printf("r:%d\r\n", r)
	r = trap([]int{5,5,1,7,1,1,5,2,7,6})
	fmt.Printf("r:%d\r\n", r)
	r = trap([]int{5,5,4,7,8,2,6,9,4,5})
	fmt.Printf("r:%d\r\n", r)
	r = trap([]int{5,3,7,7})
	fmt.Printf("r:%d\r\n", r)
	 */
}
