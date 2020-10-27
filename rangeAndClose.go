
package main

import "fmt"

func main() {

	c := make(chan int)

	go meuloop(10, c)

	for v := range c {
		fmt.Println("Recebido Canal:", v)
	}

}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}
