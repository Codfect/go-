package main

import "fmt"

var x [5]int
var y [6]float64

func main() {
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5
	
	y[0] = 10.1
	
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	
	fmt.Println(y[0])
	fmt.Println(x[1])
	
	fmt.Println(len(x))
	fmt.Println(len(y))
}
