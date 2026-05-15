package operaciones

import (
	"fmt"
)

func Sumar(a, b int) int {
	return a + b
}

func Restar(a, b int) int {
	return a - b
}

func Multiplicar(a, b int) int {
	return a * b
}

func Dividir(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: No se puede dividir por cero.")
		return 0
	}
	return a / b
}
