package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)
	defer out.Flush()
	//read_int := func() (x int) {
	//	c, _ := in.ReadByte()
	//	for ; c < '0'; c, _ = in.ReadByte() {
	//
	//	}
	//	for ; c >= '0'; c, _ = in.ReadByte() {
	//		x = x*10 + int(c-'0')
	//	}
	//	return
	//}
	//num := read_int()
	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	fmt.Fprintln(out, a, b, c)
	//fmt.Fprintln(out, num)
	return
}
