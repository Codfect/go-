package main

import (
	"fmt"
)

func main() {
	//Array é fixo e imutavel
	array := [5]int {1, 2, 3, 4, 5}
	array2 := [5]int {}
	
	//Slice flexivel e possivel alterar
	slice := []int {1, 2, 3, 4, 5}
	
	slice2 := append(slice, 6, 7, 8, 9, 10)
	
	slice[3] = 777
	
	fmt.Println(array)
	fmt.Println(array2)
	fmt.Printf("\n")
	
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice[3])
}
