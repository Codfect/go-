package main

import (
	"fmt"
)

type pessoa struct {
	nome    string
	last    string
	sabores []string
}

func main() {

	//mapa := map[string]pessoa
	
	mapa := make(map[string]pessoa)
	
	mapa["Korbes"] = pessoa{
		nome: "Ellen",
		last: "Korbes",
		sabores: []string {"pistache", "blue ice", "coconut"},
	}
	
	mapa["Jarvis"] = pessoa{
		nome: "Luscius",
		last: "Jarvis",
		sabores: []string {"oil", "gas", "scripts"},
	}
	
	for _, val := range mapa {
		fmt.Println(val)
	}
}
