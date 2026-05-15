package operaciones

import (
	"errors"
)

func Operaciones(a, b int) (int, int) {
	return a + b, a - b
}

func Dividir2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Error: No se puede dividir por cero.")
	}
	return a / b, nil
}
