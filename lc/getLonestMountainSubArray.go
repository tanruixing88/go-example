package main

func getLonestMountainSubArray(a []int) int {
	if len(a) <= 3 {
		return 0
	}

	subarrayLen := 1
	maxSubArrayLen := subarrayLen
	dir := 2 // 2 平 1 上 0 下
	lastDir := 2
	//1,2,3,4,3,2,4,5,8
	//5,4,3,2,1,1
	//1,2,1
	for i := 0; i < len(a)-1; i++ {
		if dir == 1 {
			if a[i+1] > a[i] {
				subarrayLen++
			} else if a[i+1] == a[i] {
				lastDir = 1
				dir = 2
				subarrayLen = 1
			} else {
				lastDir = 1
				dir = 0
				subarrayLen++
			}
		} else if dir == 0 {
			if a[i+1] < a[i] {
				subarrayLen++
			} else if a[i+1] == a[i] {
				if subarrayLen > maxSubArrayLen && lastDir == 1 {
					maxSubArrayLen = subarrayLen
				}
				lastDir = 0
				dir = 2
				subarrayLen = 1
			} else {
				if subarrayLen > maxSubArrayLen && lastDir == 1 {
					maxSubArrayLen = subarrayLen
				}
				lastDir = 0
				dir = 1
				subarrayLen = 2
			}
		} else {
			if a[i+1] <= a[i] {
				subarrayLen = 1
			} else {
				subarrayLen = 2
				dir = 1
				lastDir = 2
			}
		}
	}

	// 1,2,3,4,5,4
	if dir == 0 && lastDir == 1 && subarrayLen > maxSubArrayLen {
		maxSubArrayLen = subarrayLen
	}

	if maxSubArrayLen < 3 {
		return 0
	} else {
		return maxSubArrayLen
	}
}



// course['startTime'] + 14 * 86400 < curTime && role == user
// 1. go value
// 2. go lua js// for // error
// 3. gengine -> antlr4 - json-config -> json-engine

