package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	c := make([]int, 3)
	copy(c, a)
	fmt.Println(c, a)
	c[1] = 3
	fmt.Println(c, a)
}
