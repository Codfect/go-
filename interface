package main

import (
	"fmt"
)

type player struct {
	name    string
	team    string
	age     int
	posicao string
	stats
	scores
}

type stats struct {
	height    float64
	number    int
	draftyear int
}

type scores struct {
	points  float64
	rebound float64
	assists float64
}

func (v player) interview() {
	fmt.Println(
		"Me chamo", v.name, "tenho", v.age, "anos", "e jogo no", v.team, "na posição", v.posicao, ". Iniciei no draft de",
		v.draftyear, "com a camisa", v.number, ". A minha altura é de", v.height,
		". A minha média por partida é de", v.points, ", pontos,", v.rebound, "rebotes e", v.assists, "assistencias.",
	)
}

//Ambos podem acessar a função "NBA" que chama o método "interview" de cada "jogador"

type jogador interface {
	interview()
}

func NBA(j jogador) {
	j.interview()
}

func main() {
	DamianLillard := player{
		name:    "Damian Lillard",
		team:    "Portland Trail Brazzers",
		age:     30,
		posicao: "Guard (G)",
		stats: stats{
			height:    1.88,
			number:    0,
			draftyear: 2012,
		},
		scores: scores{
			points:  30.0,
			rebound: 4.3,
			assists: 8.0,
		},
	}

	JaMorant := player{
		name:    "Ja Morant",
		team:    "Memphis Grizzlies",
		age:     21,
		posicao: "Guard (G)",
		stats: stats{
			height:    1.90,
			number:    12,
			draftyear: 2019,
		},
		scores: scores{
			points:  17.8,
			rebound: 3.9,
			assists: 7.3,
		},
	}

	DamianLillard.interview()

	NBA(DamianLillard)
	fmt.Printf("\n")
	NBA(JaMorant)
}
