package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 1
		c <- 1
		fmt.Println(<-c)
	}()
	v := <-c
	fmt.Println(<-c)
	fmt.Println(v)
}
