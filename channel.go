package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	
	go func() {
		canal <- 49	
	}()
	

	fmt.Println(<-canal)
}
