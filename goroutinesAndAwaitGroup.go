package main

import (
	"fmt"
	"sync"
	"time"
)

//Package level scope, para poder ser usado dentro dos scopos das funções
var wg sync.WaitGroup

func main() {

	//add (adiciona o número de goroutines)
	wg.Add(2)

	go func1()
	go func2()

	//Espera as goroutines serem executadas primeiro
	wg.Wait()
}

func func1() {
	for i := 0; i <= 70; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	//deu! Pode voltar a função Main
	wg.Done()
}

func func2() {
	for i := 0; i <= 70; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	wg.Done()
}
