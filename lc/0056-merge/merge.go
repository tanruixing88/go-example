package main

import "fmt"

func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals) - 1 - i; j++ {
			if intervals[j+1][0] < intervals[j][0] {
				intervals[j+1], intervals[j] = intervals[j], intervals[j+1]
			}
		}
	}

	mergeIntervals := make([][]int, 0)
	if len(intervals) <= 0 {
		return mergeIntervals
	}
	curMergeInterval := append([]int{}, intervals[0]...)

	max := func(x, y int) int { if x >= y {return x} else {return y}}
	min := func(x, y int) int { if x >= y {return y} else {return x}}
	for i := 1; i < len(intervals); i++ {
		if curMergeInterval[1] >= intervals[i][0] {
			curMergeInterval = []int{min(curMergeInterval[0], intervals[i][0]), max(curMergeInterval[1], intervals[i][1])}
			continue
		}

		mergeIntervals = append(mergeIntervals, curMergeInterval)
		curMergeInterval = append([]int{}, intervals[i]...)
	}

	mergeIntervals = append(mergeIntervals, curMergeInterval)

	return mergeIntervals
}

func main() {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Printf("intervals:%+v merge:%+v\r\n", intervals, merge(intervals))

	intervals = [][]int{{1,4},{4,5}}
	fmt.Printf("intervals:%+v merge:%+v\r\n", intervals, merge(intervals))
}
