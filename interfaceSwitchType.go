package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentes  float64
	salario float64
}

type arquiteto struct {
	pessoa
	construcoes string
	tamanho     string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e Bom dia!")
}

//Ambos podem acessar a função serHumano que chama o método oibomdia de cada gente

type gente interface {
	oibomdia()
}

func serHumano(g gente) {
	g.oibomdia()
	switch g.(type){
		case dentista:
			fmt.Println("Eu ganho:", g.(dentista).salario, "reais")
			
		case arquiteto:
			fmt.Println("Eu construo:", g.(arquiteto).construcoes)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "Da Silva",
			idade:     50,
		},
		dentes:  50.110,
		salario: 10.011,
	}

	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Prédio",
			sobrenome: "Dos Santos",
		},
		construcoes: "Torres",
		tamanho:     "Mega",
	}
	
	serHumano(mrdente)
	fmt.Println("")
	serHumano(mrpredio)
}
