package main

import (
	"bufio"
	"fmt"
	"os"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			if nums[l] <= target || nums[r] >= target && nums[r] > nums[m] {
				// m和target 都在左半边
				// 或m和target 都在右半边
				r = m - 1
			} else {
				// m在左半边 target在右半边
				l = m + 1
			}
		} else {
			if nums[m] >= nums[l] || nums[r] >= target && nums[r] > nums[m] {
				// m和target 都在左半边
				// 或m和target 都在右半边
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(in, &arr[i])
	}
	fmt.Fprintln(out, search(arr, k))
	return
}
