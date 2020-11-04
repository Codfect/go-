package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	lado1 float64
	lado2 float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() {
	resultado := r.lado1 * r.lado2
	fmt.Println("A Area do retangulo é:", resultado)
}

func (c circulo) area() {
	perimetro := math.Pi * c.raio * c.raio
	fmt.Println(perimetro)
}

//Um tipo "figura" que define como interface
//qualquer tipo que tiver o método area
type figura interface {
	area()
}

//Função info que recebe um tipo "figura" e retorna
//a area da figura
func info(f figura) {
	f.area()
}

func main() {
	umretangulo := retangulo{
		lado1: 65.5,
		lado2: 74.5,
	}

	umcirculo := circulo{ 
		raio: 7.7,
	}

	umretangulo.area()
	fmt.Println("")
	umcirculo.area()
	fmt.Println("")
	
	info(umcirculo)
	info(umretangulo)
}
