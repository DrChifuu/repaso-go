package main

import (
	"fmt"
	"pruebago/calcular"
)

var message string
var a, b int = 10, 20

func main() {
	fmt.Println("hola mundo")

	message = "esto es un mensaje\n"

	name := "go "
	age := 15
	pi := 3.1416
	activo := true

	const max = 100

	fmt.Print(message, name, age, pi, activo, max, "\n")

	var x int16 = 10
	var y int16 = 20

	fmt.Println("Suma: ", calcular.Sumar(x, y))
	fmt.Println("Resta: ", calcular.Restar(x, y))
	fmt.Println("Multiplicar: ", calcular.Multiplicar(x, y))
	fmt.Println("Dividir: ", calcular.Dividir(x, y))

}
