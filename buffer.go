package main

import (
	"fmt"
)

func main() {
	canal := make(chan int, 1)
	
		canal <- 7
	
	fmt.Println(<-canal)
}
