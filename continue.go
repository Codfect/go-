package main

import "fmt"

func main() {
	
	x := 0
	for x < 20 {
		x++
		if (x % 2 == 0) { 
			continue //Quebra a iteração do loop no meio
		}
		fmt.Println(x)
	}
}
