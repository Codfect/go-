package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	profissao string
	salario   int
}

func main() {

	Funcionario1 := pessoa{
		nome:  "João",
		idade: 21,
	}

	Funcionario2 := profissional{
		pessoa: pessoa{
			nome:  "Maria",
			idade: 23,
		},
		profissao: "Pizzaiola",
		salario:   10000,
	}

	fmt.Println(Funcionario1)
	fmt.Println(Funcionario2)
}
