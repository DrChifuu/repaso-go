package main

import (
	"fmt"
	"pruebago/calcular"
)

var mensaje string
var a, b int = 10, 20

func main() {
	fmt.Println("hola mundo")

	mensaje = "caca pichi poto "

	nombre := "go "
	edad := 15
	pi := 3.1416
	activo := true

	const maximo = 100

	fmt.Print(mensaje, nombre, edad, pi, activo, maximo)

	var x int16 = 10
	var y int16 = 20

	fmt.Print(calcular.Sumar(x, y))

}
