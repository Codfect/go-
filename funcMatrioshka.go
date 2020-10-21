package main

import (
	"fmt"
)

func main() {
	
	x := matrioshka()
	
	y := x(2)
	fmt.Println(y)
	
	fmt.Println(matrioshka()(3))
}

func matrioshka() func(int) int {
	return func(variavel int) int {
		return variavel * 10
	}
}
