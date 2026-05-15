package saludos

import "fmt"

func Saludar(nombre string) string {
	return fmt.Sprintf("Hola, %s!", nombre)
}

func holaMundo() string {
	return "¡Hola, mundo!"
}
