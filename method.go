package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (variavel pessoa) nomedafunc() {
	fmt.Println(variavel.idade, variavel.nome, "diz Bom dia!")
}

func main() {
	funcionario := pessoa{
		nome: "Mauricio",
		idade: 32,
	}

	funcionario.nomedafunc()
}
