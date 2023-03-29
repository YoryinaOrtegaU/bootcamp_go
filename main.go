package main

import (
	"fmt"
	f "pratica/funciones"
)

func main() {

	fmt.Println(f.CalcularImpuesto(51000))

	fmt.Println(f.CalcularPromedio(2, 3, 5, 5, 5, 4))
	fmt.Println(f.CalcularSalario(15, "B"))

	calculo := f.CalcularEstadisticas("promedio")
	resultado := calculo(1, 2, 3, 4, 5)

	fmt.Println(resultado)

	animal := f.Animal("gato")
	fmt.Println(animal(5))

}
