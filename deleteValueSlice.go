package main

import "fmt"

func main() {
	sabores := []string{"Peperoni", "Mozzarella", "Marguerita", "Diavola", "Stagioni"}

	sabores = append(sabores[:1], sabores[3:]...)

	fmt.Println(sabores)
}
