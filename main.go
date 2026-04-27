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
	// Bienvenida := "Hola, bienvenido a Go!"
	// var nombre string
	// //var edad int
	// var numero1, numero2, resultado int

	// fmt.Println(Bienvenida)
	// fmt.Print("Por favor, ingresa tu nombre: ")
	// fmt.Scanln(&nombre)
	// fmt.Print("Por favor, ingresa el primer número 1: ")
	// fmt.Scanln(&numero1)
	// fmt.Print("Por favor, ingresa el segundo número 2: ")
	// fmt.Scanln(&numero2)
	// // fmt.Print("Por favor, ingresa tu edad: ")
	// // fmt.Scanln(&edad)
	// resultado = numero1 + numero2
	// fmt.Println("El resultado de la suma es: ", resultado)
	//fmt.Println("Hola", nombre, "tienes", edad, "años. ¡Bienvenido a Go!")
	//fmt.Println("Hello, World! 😎\n soy gigpz")
	//fmt.Println(nombre, edad, esProgramador, salario, PI)

	// a := 20
	// b := 10

	// fmt.Println("Operadores Aritméticos")
	// fmt.Println("Suma:", a+b)
	// fmt.Println("Resta:", a-b)
	// fmt.Println("Multiplicación:", a*b)
	// fmt.Println("División:", a/b)
	// fmt.Println("Módulo:", a%b) // 7 / 3 = 1, 5 / 2 = 2, 1

	// fmt.Println("Operadores Relacionales")
	// fmt.Println("a es igual a b:", a == b)
	// fmt.Println("a es diferente de b:", a != b)
	// fmt.Println("a es mayor que b:", a > b)
	// fmt.Println("a es menor que b:", a < b)
	// fmt.Println("a es mayor o igual que b:", a >= b)
	// fmt.Println("a es menor o igual que b:", a <= b)

	// fmt.Println("Operadores Lógicos")
	// fmt.Println("a es mayor que 15 y b es menor que 20:", a > 15 && b < 20) //AND
	// fmt.Println("a es mayor que 25 o b es menor que 15:", a > 25 || b < 15) //OR
	// fmt.Println("No es cierto que a sea menor que b:", !(a < b))            //NOT

	// edad := 17 //asignacion
	// tieneDni := true

	// puedeIngresar := edad >= 18 && tieneDni
	// fmt.Println("¿Puede ingresar al evento?", puedeIngresar)
	// var edad int
	// esVIP := false

	// fmt.Print("Que edad tienes ?:")
	// fmt.Scan(&edad)

	// puedeIngresar := edad >= 18 || esVIP
	// fmt.Println("¿Puede ingresar al evento?", puedeIngresar)

	/////Conversión de tipos///////

	// var numero int = 10
	// var decimal float64 = float64(numero) + 0.12
	// fmt.Println(decimal)

	//Tabla de verdad de AND
	// A     B     A AND B
	// true  true  true
	// true  false false
	// false true  false
	// false false false

	//Tabla de verdad de OR
	// A     B     A OR B
	// true  true  true
	// true  false true
	// false true  true
	// false false false

	//Tabla de verdad de NOT
	// A     NOT A
	// true  false
	// false true

	//Reto 1: Calculadora básica
	var num1, num2 int
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&num2)
	fmt.Println("Suma:", num1+num2)
	fmt.Println("Resta:", num1-num2)
	fmt.Println("Multiplicación:", num1*num2)
	fmt.Println("División Entera:", num1/num2)
	fmt.Println("División Decimal:", float64(num1)/float64(num2))
	fmt.Println("Módulo:", num1%num2)

	//Reto 2: Número par o impar
	var numero int
	fmt.Print("Ingresa un número para verificar si es par o impar: ")
	fmt.Scan(&numero)

	esPar := numero%2 == 0
	fmt.Println("¿Es un número par?", esPar)

	//Reto 3: Validación de acceso (realista) Edad ≥ 18 && tieneDNI == true || esVIP == true
	var edad int
	var tieneDNI bool
	var esVIP bool
	fmt.Print("Ingresa tu edad: ")
	fmt.Scan(&edad)
	fmt.Print("¿Tienes DNI? (true/false): ")
	fmt.Scan(&tieneDNI)
	fmt.Print("¿Eres VIP? (true/false): ")
	fmt.Scan(&esVIP)
	accesoPermitido := (edad >= 18 && tieneDNI) || esVIP
	fmt.Println("¿Acceso permitido?", accesoPermitido)

	//Reto 4: Reto 4: Simula un login simple: usuario == "admin" && password == "1234"
	usuario := "admin"
	password := "1234"
	loginExitoso := usuario == "admin" && password == "1234"
	fmt.Println("¿Login exitoso?", loginExitoso)

	//Reto 5: Mayor de 3 números (Pide 3 números y muestra el mayor usando operadores.)
	var a, b, c int
	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&b)
	fmt.Print("Ingresa el tercer número: ")
	fmt.Scan(&c)

	mayorNumero := (a > b && a > c) || (b > a && b > c) || (c > a && c > b)
	fmt.Println("¿Cuál es el mayor?", mayorNumero)

}
