package main

import (
	"bufio"
	"fmt"
	"os"
)

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return true
		}
		if nums[l] == nums[m] && nums[m] == nums[r] {
			r--
			continue
		}
		if nums[m] >= nums[l] {
			// m在左边
			if nums[l] > target || nums[m] < target {
				// target 在右边 或 target在左边但大于m
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			// m在右边
			if nums[r] < target || nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return false
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
