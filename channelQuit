package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)
	
	go recebeQuit(canal, quit)
	enviaCanal(canal, quit)

}

func recebeQuit(channel chan int, quit chan int) {
	for i := 0; i <= 70; i++ {
		fmt.Println("Recebido", <-channel)
	}
	quit <- 0 //Quando o laço termina envia um valor
		  //Que encerra o loop infinito em case <-quit
}

func enviaCanal(channel chan int, quit chan int) {
	v := 7
	for {
		select{
			case channel <- v:  //enviando para
				v++	    //recebeQuit()	
			case <-quit:
				return
		}
	}
}
