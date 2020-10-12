package main

import "fmt"

func main() {

	for horas := 0; horas <= 12; horas ++ {
		fmt.Print("Hora: ", horas, "\n")
		
		for minutos := 0; minutos <= 60; minutos ++ {
			fmt.Print(" ", minutos)	
		}
		fmt.Print("\n\n")
	}
}
