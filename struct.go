package main

import (
	"fmt"
)

type players struct {
	nome       string
	time       string
	aposentado bool
}

func main() {

	player1 := players {
		nome:       "Julius Irving",
		time:       "Philadelphia 76s",
		aposentado: true,
	}
	
	player2 := players {nome: "Ja Morant", time: "Grizzlies", aposentado: false}

	fmt.Println(player1)
	fmt.Println(player2)
}
