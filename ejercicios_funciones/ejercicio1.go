package main

import "fmt"

/* Una empresa marinera necesita calcular el salario de sus empleados basándose en la 
cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos 
trabajados por mes, la categoría y que devuelva su salario.
 */

func CalcularSalario(minutos int, categoria string) float64{
	var horasTrabajadas float64
	horasTrabajadas = float64(minutos)/60

	var salario float64

	switch categoria{
	case "C":
		salario := 1000*horasTrabajadas
	case "B":
		salario := 1500*horasTrabajadas
	case "A":
		salario := 3000*horasTrabajadas
	default:
		

	}
	return salario
}

func CalcularPromedio(calificaciones ...float64) (float64, error){
	var promedio float64

	for _, calificacion := range calificaciones{

		promedio += calificacion
	}
	return promedio/float64(len(calificaciones))
}

func CalcularImpuesto(salario float64) float64{
	var impuesto float64
	if salario > 150000{
		impuesto = salario*0.27	
	} else if salario >  50000{
		impuesto = salario*0.17
	}
	return impuesto
}
 
func main(){
	fmt.Println(CalcularImpuesto(51000))

	fmt.Println(CalcularPromedio(5, 6, 9, 10))
}