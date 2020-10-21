package main

import (
	"encoding/json"
	"os"
)

type aventureiro struct {
	Nome       string
	Raca       string
	Habilidade string
	Classe     string
	Arma       string
}

func main() {

	ufgar := aventureiro{
		Nome:       "Ufgar",
		Raca:       "Anão",
		Habilidade: "Proficiência com ferramentas",
		Classe:     "Guerreiro",
		Arma:       "Machado",
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(ufgar)
	encoder.Encode(ufgar.Raca)
}
