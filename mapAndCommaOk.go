package main

import (
	"fmt"
)

func main() {

	amigos := map[string]int{
		"Alfredo": 123123,
		"Maria":   321321,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["Maria"])

	amigos["Gopher"] = 100100

	fmt.Println(amigos)
	fmt.Println(amigos["Gopher"])

	//comma ok idiom
	if sera, ok := amigos["fantasma"]; !ok {
		fmt.Println("Não tem")
	} else {
		fmt.Println(sera, ok)
	}

}
