package main

import "fmt"

func main() {
	quemestanoescritorio := "Jaco"

	switch quemestanoescritorio {
		case "João", "Jaco":
			fmt.Println("João ou Jaco estão no escritorio")
		case "Marquinhos":
			fmt.Println("Marquinhos está no escritorio")
			fallthrough
		case "Maria":
			fmt.Println("Quando Marquinhos vem ao escritorio, Maria também vem")
		case "Joana":
			fmt.Println("Joana está no escritorio")
		default:
			fmt.Println("Não tem ninguém no escritorio")
	}
}
