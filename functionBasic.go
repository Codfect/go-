package main

import "fmt"

func main() {
	//Ao criar uma função o nome é parametro
	//Quando ta chamando a função o nome é argumento

	basica()
	bom("manhã")
	bom("tarde")
	bom("qualquercoisa")
}

func basica() {
	fmt.Println("Hello Basic")
}

func bom(s string) {
	if s == "manhã" {
		fmt.Println("Bom dia!")
	} else if s == "tarde" {
		fmt.Println("Boa tarde!")
	} else {
		fmt.Println("Boa noite!")
	}
}
