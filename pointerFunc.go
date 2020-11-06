package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	ze := pessoa {
		nome:      "Zezinho",
		sobrenome: "da Silva",
		idade:     10,
	}

	fmt.Println(ze)
	mudeMe(&ze)
	fmt.Println(ze)
}

func mudeMe(variavel *pessoa) {
	(*variavel).nome = "Joaozinho"
}
