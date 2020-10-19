package main

import (
	"fmt"
)

func main() {
	fmt.Println(fatorial(7))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}
