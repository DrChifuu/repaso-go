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

	/*
		uso de len para deteccion de cantidad de elementos de un slice
	*/

	lista_slice := []int{10, 20, 30}
	fmt.Println("cantidad datos en la lista", len(lista_slice))

	numeros_cap := make([]int, 0, 5)

	fmt.Println(numeros_cap)

	for i := 0; i < 5; i++ {
		numeros_cap = append(numeros_cap, i+1)
	}

	fmt.Println("Post uso de espacios pre asignados ", numeros_cap)

	/*
		Si tuvieramos la situación de valores del slice ocupados
	*/

	numeros_cap2 := make([]int, 3, 5)

	fmt.Println("lista pre ocupada con 0 en 3 espacios del sub espacio 5", numeros_cap2)

	for i := 0; i < 3; i++ {
		numeros_cap2[i] = (i + 1)
	}

	fmt.Println("Post uso de espacios pre asignados ", numeros_cap2)

	numeros_cap2 = append(numeros_cap2, 3)
	numeros_cap2 = append(numeros_cap2, 5)

	fmt.Println(numeros_cap2)

	type Genero string

	const (
		Hombre Genero = "hombre"
		Mujer  Genero = "mujer"
		Goblin Genero = "goblin"
	)

	type Personaje struct {
		Nombre string
		Vida   float32
		Genero Genero
	}

	p := Personaje{Nombre: "pepito", Vida: 66.7, Genero: "goblin"}

	print("EL genero del personaje 'p' es: ", p.Genero)

}
