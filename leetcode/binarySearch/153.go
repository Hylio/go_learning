package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMin(nums []int) int {
	// 数组是部分有序的 分为两个递增序列 可以使用二分的方法
	// 如果 nums[m] > nums[r]
	// 说明m在左边的递增序列 最小值一定在m的右侧
	// 反之则说明m在右边的递增序列 最小值一定是m 或在m的左侧
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(in, &arr[i])
	}
	fmt.Fprintln(out, findMin(arr))
	return
}
