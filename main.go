package main

import (
	"fmt"
)

func saludar(nombre string) {
	fmt.Println("Hola", nombre, "bienvenido a Go!")
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: No se puede dividir por cero.")
		return 0
	}
	return a / b
}

func operaciones(a, b int) (int, int) {
	return a + b, a - b
}

func dividir2(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Error: No se puede dividir por cero."
	}
	return a / b, "División exitosa"
}

var x int = 10
var y int = 20

func inicial() {
	y = 0
}

func prueba() {
	x := 15 + y
	fmt.Println("Valor de x:", x)
}

// Retos:
// -Convertidor de moneda: solesADolares(monto float64) float64
func solesADolares(monto float64) float64 {
	const tasaCambio = 3.51 // Ejemplo de tasa de cambio
	return monto * tasaCambio
}

// -Sistema escolar: evaluarNota(nota int) string
func evaluarNota(nota int) string {
	if nota >= 18 {
		return "Excelente"
	} else if nota >= 14 {
		return "Muy Bien"
	} else if nota >= 11 {
		return "Bien"
	} else {
		return "Insuficiente"
	}
}

// -Login reutilizable: login(usuario string, password string) bool
func login(usuario string, password string) bool {
	return usuario == "admin" && password == "1234"
}

// MINI PROYECTO DEL MÓDULO
//  Sistema bancario modular

// Menú:

// 1. Consultar saldo
// 2. Depositar
// 3. Retirar
// 4. Salir

var saldo float64 = 0.00

func menu() {
	for {
		fmt.Println("============Bienvenido al Cajero Automatico===============")
		fmt.Println("1. Ver saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Retirar")
		fmt.Println("4. Salir")
		fmt.Print("Selecciona una opción: ")
		var opcion int
		fmt.Scan(&opcion)
		switch opcion {
		case 1:
			verSaldo()
		case 2:
			depositar()
		case 3:
			retirar()
		case 4:
			fmt.Println("Gracias por usar el Cajero Automático. ¡Hasta luego!")
			return
		default:
			fmt.Println("Opción inválida. Por favor, selecciona una opción válida.")
		}
	}
}

func depositar() {
	fmt.Print("Ingrese el monto a depositar: ")
	var deposito float64
	fmt.Scan(&deposito)
	saldo += deposito
	fmt.Println("Depósito exitoso. Nuevo saldo:", saldo)
}

func retirar() {
	fmt.Print("Ingrese el monto a retirar: ")
	var retiro float64
	fmt.Scan(&retiro)
	if retiro > saldo {
		fmt.Println("Saldo insuficiente.")
	} else {
		saldo -= retiro
		fmt.Println("Retiro exitoso. Nuevo saldo:", saldo)
	}
}

func verSaldo() {
	fmt.Println("Tu saldo actual es: $", saldo)
}

func main() {

	saludar("cristhian")
	menu()
	fmt.Println("====================fin del programa==========================")
	// fmt.Println("Suma 1:", sumar(10, 20))

	// fmt.Println("Resta 1:", restar(10, 5))

	// fmt.Println("Multiplicación 1:", multiplicar(5, 5))

	// fmt.Println("División 1:", dividir(10.8, 2))
	// fmt.Println("División 2:", dividir(20, 0))

	// suma, _ := operaciones(15, 5)
	// fmt.Println("Suma 2:", suma)
	// // fmt.Println("Resta 2:", resta)

	// _, mensaje := dividir2(20, 5)
	// //fmt.Println("División 3:", resultado)
	// fmt.Println(mensaje)

	// prueba()
	// fmt.Println("Conversión de soles a dólares:", solesADolares(100))
	// fmt.Println("Evaluación de nota 18:", evaluarNota(18))
	// fmt.Println("Evaluación de nota 16:", evaluarNota(16))
	// fmt.Println("Evaluación de nota 12:", evaluarNota(12))
	// fmt.Println("Evaluación de nota 10:", evaluarNota(10))
	// fmt.Println("Login con usuario 'admin' y contraseña '1234':", login("admin", "1234"))
	// fmt.Println("Login con usuario 'user' y contraseña 'abcd':", login("user", "abcd"))

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

	// //Reto 1: Calculadora básica
	// var num1, num2 int
	// fmt.Print("Ingresa el primer número: ")
	// fmt.Scan(&num1)
	// fmt.Print("Ingresa el segundo número: ")
	// fmt.Scan(&num2)
	// fmt.Println("Suma:", num1+num2)
	// fmt.Println("Resta:", num1-num2)
	// fmt.Println("Multiplicación:", num1*num2)
	// fmt.Println("División Entera:", num1/num2)
	// fmt.Println("División Decimal:", float64(num1)/float64(num2))
	// fmt.Println("Módulo:", num1%num2)

	// //Reto 2: Número par o impar
	// var numero int
	// fmt.Print("Ingresa un número para verificar si es par o impar: ")
	// fmt.Scan(&numero)

	// esPar := numero%2 == 0
	// fmt.Println("¿Es un número par?", esPar)

	// //Reto 3: Validación de acceso (realista) Edad ≥ 18 && tieneDNI == true || esVIP == true
	// var edad int
	// var tieneDNI bool
	// var esVIP bool
	// fmt.Print("Ingresa tu edad: ")
	// fmt.Scan(&edad)
	// fmt.Print("¿Tienes DNI? (true/false): ")
	// fmt.Scan(&tieneDNI)
	// fmt.Print("¿Eres VIP? (true/false): ")
	// fmt.Scan(&esVIP)
	// accesoPermitido := (edad >= 18 && tieneDNI) || esVIP
	// fmt.Println("¿Acceso permitido?", accesoPermitido)

	// //Reto 4: Reto 4: Simula un login simple: usuario == "admin" && password == "1234"
	// usuario := "admin"
	// password := "1234"
	// loginExitoso := usuario == "admin" && password == "1234"
	// fmt.Println("¿Login exitoso?", loginExitoso)

	// //Reto 5: Mayor de 3 números (Pide 3 números y muestra el mayor usando operadores.)
	// var a, b, c int
	// fmt.Print("Ingresa el primer número: ")
	// fmt.Scan(&a)
	// fmt.Print("Ingresa el segundo número: ")
	// fmt.Scan(&b)
	// fmt.Print("Ingresa el tercer número: ")
	// fmt.Scan(&c)

	// mayorNumero := (a > b && a > c) || (b > a && b > c) || (c > a && c > b)
	// fmt.Println("¿Cuál es el mayor?", mayorNumero)

	/////////Módulo 4: Estructuras de Control en Go/////////
	//1 Condicionales: if, else if, else

	// var edad int
	// fmt.Print("Ingresa tu edad: ")
	// fmt.Scan(&edad)

	// if edad <= 12 {
	// 	fmt.Println("Eres un niño.")
	// } else if edad <= 18 {
	// 	fmt.Println("Eres un adolescente.")
	// } else {
	// 	fmt.Println("Eres mayor de edad.")
	// }

	// //Reto 5: Mayor de 3 números (Pide 3 números y muestra el mayor usando operadores.)
	// var a, b, c int
	// fmt.Print("Ingresa el primer número: ")
	// fmt.Scan(&a)
	// fmt.Print("Ingresa el segundo número: ")
	// fmt.Scan(&b)
	// fmt.Print("Ingresa el tercer número: ")
	// fmt.Scan(&c)

	// if a > b && a > c {
	// 	fmt.Println("El mayor es a:", a)
	// } else if b > a && b > c {
	// 	fmt.Println("El mayor es b:", b)
	// } else {
	// 	fmt.Println("El mayor es: c:", c)
	// }

	// //2 switch
	// var dia int
	// fmt.Print("Ingresa un número del 1 al 7 para conocer el día de la semana: ")
	// fmt.Scan(&dia)

	// switch dia {
	// case 1:
	// 	fmt.Println("Lunes")
	// case 2:
	// 	fmt.Println("Martes")
	// case 3:
	// 	fmt.Println("Miércoles")
	// case 4:
	// 	fmt.Println("Jueves")
	// case 5:
	// 	fmt.Println("Viernes")
	// case 6:
	// 	fmt.Println("Sábado")
	// case 7:
	// 	fmt.Println("Domingo")
	// default:
	// 	fmt.Println("Número inválido")
	// }

	// edad2 := 17
	// switch {
	// case edad2 < 13:
	// 	fmt.Println("Eres un niño.")
	// case edad2 < 18:
	// 	fmt.Println("Eres un adolescente.")
	// default:
	// 	fmt.Println("Eres mayor de edad.")
	// }

	//2 Ciclos o repetitivos: for, while (simulado con for), do-while (simulado con for)
	//for → cuando sabes cuántas veces repetir
	//while → cuando repites mientras se cumpla una condición

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Iteración número:", i)
	// }

	// contador := 0
	// for contador < 5 {
	// 	fmt.Println("Contador:", contador)
	// 	contador += 1
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Número:", i)
	// 	if i == 5 {
	// 		break //sale del ciclo
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 != 0 {
	// 		continue //salta a la siguiente iteración
	// 	}
	// 	fmt.Println("Número par:", i)
	// }
	//infinito
	// for {
	// 	fmt.Println("Hola")
	// }

	// var numero, limite int
	// fmt.Print("Ingresa un número para mostrar su tabla de multiplicar: ")
	// fmt.Scan(&numero)
	// fmt.Print("Ingrese numero limite de multiplicacion: ")
	// fmt.Scan(&limite)

	// for i := 1; i <= limite; i++ {
	// 	fmt.Println("multipliacion de", numero, "x", i, "=", numero*i)
	// }
	// fmt.Println("==============================================")
	// // Reto 1:Login con 3 intentos admin, 1234 - usar for
	// fmt.Println("Bienvenido a el sistema ingrese sus credenciales")
	// var usuario, password string
	// for intentos := 1; intentos <= 3; intentos++ {
	// 	fmt.Print("Usuario: ")
	// 	fmt.Scan(&usuario)
	// 	fmt.Print("Contraseña: ")
	// 	fmt.Scan(&password)

	// 	if usuario == "admin" && password == "1234" {
	// 		fmt.Println("Login exitoso")
	// 		break
	// 	} else {
	// 		fmt.Println("Credenciales incorrectas, intento", intentos, "de 3.")
	// 	}
	// }
	// fmt.Println("==============================================")

	// Reto 2: Menú interactivo
	// // 1. Saludar
	// // 2. Mostrar hora
	// // 3. Salir
	// // Usar for infinito + switch|if.
	// for {
	// 	fmt.Println("1. Saludar")
	// 	fmt.Println("2. Mostrar hora")
	// 	fmt.Println("3. Salir")
	// 	fmt.Print("Selecciona una opción: ")
	// 	var opcion int
	// 	fmt.Scan(&opcion)

	// 	if opcion == 1 {
	// 		fmt.Println("Hola!")
	// 	} else if opcion == 2 {
	// 		fmt.Println("La hora actual es:", time.Now().Format("15:04:05"))
	// 	} else if opcion == 3 {
	// 		fmt.Println("Saliendo...")
	// 		break
	// 	} else {
	// 		fmt.Println("Opción inválida")
	// 	}
	// 	// switch opcion {
	// 	// case 1:
	// 	// 	fmt.Println("Hola!")
	// 	// case 2:
	// 	// 	fmt.Println("La hora actual es:", time.Now().Format("15:04:05"))
	// 	// case 3:
	// 	// 	fmt.Println("Saliendo...")
	// 	// 	return
	// 	// default:
	// 	// 	fmt.Println("Opción inválida")
	// 	// }
	// }

	// fmt.Println("==============================================")
	// Reto 3: Suma acumulada
	// // Pedir números hasta que el usuario escriba 0.
	// // Luego mostrar total sumado.
	// fmt.Println("Ingrese números para sumar (0 para terminar):")
	// var suma, numero2 int
	// for {
	// 	fmt.Print("Número: ")
	// 	fmt.Scan(&numero2)
	// 	if numero2 == 0 {
	// 		break
	// 	}
	// 	suma += numero2
	// }
	// fmt.Println("Total sumado:", suma)

	// fmt.Println("==============================================")

	// Mini Proyecto del módulo
	// Sistema de cajero simple
	// 1. Ver saldo
	// 2. Depositar
	// 3. Retirar
	// 4. Salir

	// Usar:
	// for
	// switch
	// if

	// var saldo float64 = 0.00
	// for {
	// 	fmt.Println("============Bienvenido al Cajero Automatico===============")
	// 	fmt.Println("1. Ver saldo")
	// 	fmt.Println("2. Depositar")
	// 	fmt.Println("3. Retirar")
	// 	fmt.Println("4. Salir")
	// 	fmt.Print("Selecciona una opción: ")

	// 	var opcion int
	// 	fmt.Scan(&opcion)

	// 	switch opcion {
	// 	case 1:
	// 		fmt.Println("Tu saldo actual es: $", saldo)
	// 	case 2:
	// 		var deposito float64
	// 		fmt.Print("Ingrese el monto a depositar: ")
	// 		fmt.Scan(&deposito)
	// 		saldo += deposito
	// 		fmt.Println("Depósito exitoso. Nuevo saldo:", saldo)
	// 	case 3:
	// 		var retiro float64
	// 		fmt.Print("Ingrese el monto a retirar: ")
	// 		fmt.Scan(&retiro)
	// 		if retiro > saldo {
	// 			fmt.Println("Saldo insuficiente.")
	// 		} else {
	// 			saldo -= retiro
	// 			fmt.Println("Retiro exitoso. Nuevo saldo:", saldo)
	// 		}
	// 	case 4:
	// 		fmt.Println("Saliendo...")
	// 		return
	// 	default:
	// 		fmt.Println("Opción inválida")
	// 	}

	// }
}
