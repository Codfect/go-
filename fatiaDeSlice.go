package main

import "fmt"

func main() {
	sabores := []string {"Peperoni", "Mozzarella", "Marguerita", "Diavola", "Stagioni"}
	
	// [] [a:] [:b] [a:b]
	fatia := sabores [:]
	
	fmt.Println(fatia)
	
	fmt.Printf("%v, %v, %v \n\n", sabores[0], sabores[1], sabores[2])
	
	for x := 0; x < len(sabores); x ++ {
	
		fmt.Println(x, sabores[x])	
	}

}
