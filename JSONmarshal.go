package main

import (
	"encoding/json"
	"fmt"
)

type aventureiro struct {
	Nome       string
	Raça       string
	Habilidade string
	Classe     string
	Arma       string
}

func main() {

	anão := aventureiro{
		Nome:       "Ufgar",
		Raça:       "Anão",
		Habilidade: "Proficiência com ferramentas",
		Classe:     "Guerreiro",
		Arma:       "Machado",
	}

	elfo := aventureiro{
		Nome:       "Laeanna",
		Raça:       "Elfo",
		Habilidade: "Sentidos aguçados",
		Classe:     "Mago",
		Arma:       "Grimório",
	}

	halfling := aventureiro{"Tobi", "Halfling", "Furtividade natural", "Ladino", "Adaga"}

	a, err := json.Marshal(anão)
	if err != nil {
		fmt.Println(err)
	}

	e, err := json.Marshal(elfo)
	if err != nil {
		fmt.Println(err)
	}

	h, err := json.Marshal(halfling)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(a))
	fmt.Println(string(e))
	fmt.Println(string(h))
}
