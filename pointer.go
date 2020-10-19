package main

import (
	"fmt"
)

func main() {
	x := 0
	
	recebeovalor(x)
	fmt.Println(x)
	
	fmt.Println("")
	
	recebeopointer(&x)	
	fmt.Println(x)
}

func recebeovalor(x int) {
	x++
	fmt.Println("Na função sem apontar", x)
}

func recebeopointer(x *int) {
	*x++
	fmt.Println("Na função apontando:", *x)
}
