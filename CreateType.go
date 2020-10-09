package main

import "fmt"

type tipotle int

func main() {
	var x tipotle
	fmt.Printf("%v %T \n", x, x) 
	
	x = 42
	fmt.Printf("%v %T", x, x)	
}
