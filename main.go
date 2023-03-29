package main

import (
	"errors"
	"fmt"
	err "pratica/errorsp"
	f "pratica/funciones"
)

func main() {

	fmt.Println(f.CalcularImpuesto(160000))

	fmt.Println(f.CalcularPromedio(2, 3, 5, 5, 5, 4))

	fmt.Println(f.CalcularSalario(50, "C"))

	calculo := f.CalcularEstadisticas("promedio")
	resultado := calculo(1, 2, 3, 4, 5)

	fmt.Println(resultado)

	animal := f.Animal("gato")
	fmt.Println(animal(5))

	//errors

	var salary int
	salary = 120000
	fmt.Println(err.CalcularImpuesto(salary))

	e := CalcularImpuesto2(salary)
	verificarError := errors.Is(e, error1)
	fmt.Println(verificarError)

}

var error1 = errors.New("el salario es menor a 10.000")

func CalcularImpuesto2(salary int) error {
	if salary <= 10000 {
		return error1
	}
	return nil
}
