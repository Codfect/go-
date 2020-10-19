package main

import (
	"fmt"
)

func main() {
	fmt.Println(fatorial(7))
	
	fmt.Println(loop(7))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loop(x int) int {
	total := 1
	
	for x > 1 {
		total *= x
		x--
	}
	return total
}
