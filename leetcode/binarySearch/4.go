package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 此题可以扩展到x个有序数组中的第k大数字
	// 基本的原理是 为了寻找两个数组中的第k大数字
	// 比较两个数组中的第[k/2]个数字 如果某个数组中的数字较小 说明直接舍弃该数组中的前[k/2]个数字 并将k减半
	// 需要注意的是边界条件的设置
	// 1. 数组中的数字个数可能小于k/2 此时该数组直接取最后一个数字即可
	// 2. 当k等于0时 说明此时已经不需要再迭代了 返回两个数组较小的第一个数字
	n1, n2 := len(nums1), len(nums2)
	tot := n1 + n2
	if tot%2 == 1 {
		return float64(kthelement(nums1, nums2, tot/2+1))
	}
	return float64(kthelement(nums1, nums2, tot/2)+kthelement(nums1, nums2, tot/2+1)) / 2.0
}

func kthelement(nums1, nums2 []int, k int) int {
	i1, i2 := 0, 0
	for {
		if i1 == len(nums1) {
			return nums2[i2+k-1]
		}
		if i2 == len(nums2) {
			return nums1[i1+k-1]
		}
		if k == 1 {
			return min(nums1[i1], nums2[i2])
		}
		m := k / 2
		j1, j2 := min(i1+m, len(nums1))-1, min(i2+m, len(nums2))-1
		if nums1[j1] < nums2[j2] {
			k -= j1 - i1 + 1
			i1 = j1 + 1
		} else {
			k -= j2 - i2 + 1
			i2 = j2 + 1
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n1, n2 int
	fmt.Fscan(in, &n1, &n2)
	nums1 := make([]int, n1)
	nums2 := make([]int, n2)
	for i := range nums1 {
		fmt.Fscan(in, &nums1[i])
	}
	for i := range nums2 {
		fmt.Fscan(in, &nums2[i])
	}
	fmt.Fprintln(out, findMedianSortedArrays(nums1, nums2))
	return
}
