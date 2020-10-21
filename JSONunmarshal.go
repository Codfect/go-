package main

import (
	"encoding/json"
	"fmt"
)

type aventureiro struct {
	Nome       string `json:"Nome"`
	Raca       string `json:"Raca"`
	Habilidade string `json:"Habilidade"`
	Classe     string `json:"Classe"`
	Arma       string `json:"Arma"`
}

func main() {

	sb := []byte(`{"Nome":"Ufgar","Raça":"Anão","Habilidade":"Proficiência com ferramentas","Classe":"Guerreiro","Arma":"Machado"}`)

	var ufgar aventureiro
	err := json.Unmarshal(sb, &ufgar)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(ufgar)
	fmt.Println(ufgar.Arma)
}
