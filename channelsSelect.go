package main

import (
	"fmt"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	x := 100

	go func(a int) {
		for i := 0; i < a; i++ {
			canal1 <- i
		}
	}(x / 2)

	go func(a int) {
		for i := 0; i < a; i++ {
			canal2 <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-canal1:
			fmt.Println("CANAL 1:", v)
		case v := <-canal2:
			fmt.Println("CANAL 2:", v)
		}
	}
}
