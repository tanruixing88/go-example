package main

import "fmt"

func findKth(nums1 []int, nums2 []int, k int) int {
	m := len(nums1)
	n := len(nums2)

	if m + n <= 0 || k <= 0 || k > m + n {
		return -1 //异常情况
	}
	//fmt.Printf("nums1:%+v nums2:%+v k:%d\r\n", nums1, nums2, k)

	if k == 1 {
		if m == 0 && n > 0 {
			return nums2[0]
		} else if m > 0 && n == 0 {
			return nums1[0]
		} else {
			if nums1[0] < nums2[0] {
				return nums1[0]
			} else {
				return nums2[0]
			}
		}
	}

	if m == 0 {
		return nums2[k-1]
	}

	if n == 0 {
		return nums1[k-1]
	}

	mi := k/2 - 1
	ni := k/2 - 1

	if mi > m - 1 { // ni 肯定大于 n - 1
		mi = m - 1
	}

	if ni > n - 1 {
		ni = n - 1
	}

	if nums1[mi] > nums2[ni] {
		return findKth(nums1, nums2[ni+1:], k - ni - 1)
	} else if nums1[mi] < nums2[ni] {
		return findKth(nums1[mi+1:], nums2, k - mi - 1)
	} else {
		if k % 2 == 0 {
			if mi == m - 1 {
				return nums2[ni+k/2-1-(m-1)]
			}

			if ni == n - 1 {
				return nums1[mi+k/2-1-(n-1)]
			}

			return nums1[mi]
		} else {
			if mi == m - 1 {
				return nums2[ni+k/2-1-(m-1)+1]
			}

			if ni == n - 1 {
				return nums1[mi+k/2-1-(n-1)+1]
			}

			if nums1[mi+1] < nums2[ni+1] {
				return nums1[mi+1]
			} else {
				return nums2[ni+1]
			}
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	//m+n是奇数的中位数对应数组下标 (m+n)/2或者 (m+n-1)/2
	//m+n是偶数的中位数对应数组下标 (m+n)/2和(m+n)/2-1的平均值，第二个值可以转化为(m+n-1)/2，此时和m+n为奇数的情况就统一起来
	return float64(findKth(nums1, nums2, (m+n)/2+1) + findKth(nums1, nums2, (m+n-1)/2+1)) / 2.0
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{1, 2, 3}
	nums2 = []int{}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{}
	nums2 = []int{1, 2, 3}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{1}
	nums2 = []int{2}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{1,3}
	nums2 = []int{2}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{1,2}
	nums2 = []int{3,4}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{0,0}
	nums2 = []int{0,0}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{1,2,2}
	nums2 = []int{1,2,3}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)

	nums1 = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22}
	nums2 = []int{0,6}
	median = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums1:%+v nums2:%+v median:%f \r\n", nums1, nums2, median)
}
