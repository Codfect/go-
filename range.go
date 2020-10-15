package main

import (
	"fmt"
)

func main() {
	slice := []string {"banana", "maça", "morango"}
	
	for i, v := range slice {
		fmt.Println("No indicie", i, "temos o valor:", v)
	}
}
