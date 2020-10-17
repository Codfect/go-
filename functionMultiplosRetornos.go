package main

import (
	"fmt"
)

func main() {
	total, quantos := soma(7, 7, 10, 10, 10, 30, 3)

	fmt.Println(total, quantos)
}

func soma(x ...int) (int, int) {
	soma := 0

	for _, val := range x {
		soma += val
	}
	return soma, len(x)
}
