package main

import "fmt"

type hotdog int
var b hotdog

func main() {
	b = 7
	fmt.Printf("%v %T \n", b, b)

	var x =  10.0
	fmt.Printf("%v %T \n", x, x)
	
	b = hotdog(x)
	fmt.Printf("%v %T", b, b)
}
