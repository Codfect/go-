package main

import "fmt"

func main() {
	slice := []int {7, 7, 7, 7, 7, 7, 7}
	
	total := 0
	
	for _, val := range slice {
		total += val
	}
	
	fmt.Println(total)
}
