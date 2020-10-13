package main

import (
	"fmt"
)

var x interface{}

func main() {
	x = 77.7
	
	switch x.(type){
		case int:
			fmt.Printf("%v = %T", x, x)
		case float64:
			fmt.Printf("%v = %T", x, x)
		case bool:
			fmt.Printf("%v = %T", x, x)
		case string:
			fmt.Printf("%v = %T", x, x)
	}
}
