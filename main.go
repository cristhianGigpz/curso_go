package main

import "fmt"

func main() {
	// var nombre string = "Gustavo"
	// var edad int = 30
	// var esProgramador bool = true
	// var salario float64 = 300.70

	// nombre := "Gustavo"
	// edad := 30
	// esProgramador := true
	// salario := 300.70
	// const PI = 3.14159
	Bienvenida := "Hola, bienvenido a Go!"
	var nombre string
	//var edad int
	var numero1, numero2, resultado int

	fmt.Println(Bienvenida)
	fmt.Print("Por favor, ingresa tu nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Por favor, ingresa el primer número 1: ")
	fmt.Scanln(&numero1)
	fmt.Print("Por favor, ingresa el segundo número 2: ")
	fmt.Scanln(&numero2)
	// fmt.Print("Por favor, ingresa tu edad: ")
	// fmt.Scanln(&edad)
	resultado = numero1 + numero2
	fmt.Println("El resultado de la suma es: ", resultado)
	//fmt.Println("Hola", nombre, "tienes", edad, "años. ¡Bienvenido a Go!")
	//fmt.Println("Hello, World! 😎\n soy gigpz")
	//fmt.Println(nombre, edad, esProgramador, salario, PI)

}
