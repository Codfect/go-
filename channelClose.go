package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	
	go forchannels(10, canal)
	
	for v := range canal {
		fmt.Println(v)
	}
}

func forchannels(t int, c chan<- int) {
	for i := 0; i <= t; i++ {
		c <- i
	}
	close(c)
}
