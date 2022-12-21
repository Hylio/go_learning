package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMin(nums []int) int {
	// 总体思路与153大致相同
	// 但需要注意的是 如果l,r,m对应的值均相等 r会直接跳过最小值 例如[5,5,5,1,5]
	// 所以当nums[r] == nums[m]时，需要调整思路 让r--即可
	// 如果最小值在[m,r]区间，那么r总是会递减到一个更小的值，将原问题转化为153
	// 如果最小值在[l,m]区间，那么m总是会随r递减而递减到一个更小的值，从而在下一次循环中，将原问题转化为153
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			r--
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
