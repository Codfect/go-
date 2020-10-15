package main

import (
	"fmt"
)

func main() {
	caminhopara := map[int]string{
		01: "Saindo de casa",
		02: "Aguardando",
		03: "A caminho",
		04: "Chegada ao local",
	}

	fmt.Println(caminhopara)

	for indice, valor := range caminhopara{
		fmt.Println(indice, valor)
	}
	
	//Deletando a chave e valor 01
	delete(caminhopara, 01)
	
	fmt.Println(caminhopara)
}
