package main

import (
	"fmt"
	"pruebago/calcular"
	methods "pruebago/metodos"
)

var message string
var a, b int = 10, 20

func main() {

	fmt.Println("hola mundo")

	message = "esto es un mensaje\n"

	name := "go "
	age := 15
	pi := 3.1416
	active := true

	const max = 100

	fmt.Print(message, name, age, pi, active, max, "\n")

	var x int16 = 10
	var y int16 = 20

	fmt.Println("Suma: ", calcular.Sumar(x, y))
	fmt.Println("Resta: ", calcular.Restar(x, y))
	fmt.Println("Multiplicar: ", calcular.Multiplicar(x, y))
	fmt.Println("Dividir: ", calcular.Dividir(x, y))

	/*
		uso de len para deteccion de cantidad de elementos de un slice
	*/

	list_slice := []int{10, 20, 30}
	fmt.Println("cantidad datos en la lista", len(list_slice))

	numeros_cap := make([]int, 0, 5)

	fmt.Println(numeros_cap)

	for i := 0; i < 5; i++ {
		numeros_cap = append(numeros_cap, i+1)
	}

	fmt.Println("Post uso de espacios pre asignados ", numeros_cap)

	/*
		Si tuvieramos la situación de valores del slice ocupados
	*/

	number_cap2 := make([]int, 3, 5)

	fmt.Println("lista pre ocupada con 0 en 3 espacios del sub espacio 5", number_cap2)

	for i := 0; i < 3; i++ {
		number_cap2[i] = (i + 1)
	}

	fmt.Println("Post uso de espacios pre asignados ", number_cap2)

	number_cap2 = append(number_cap2, 3)
	number_cap2 = append(number_cap2, 5)

	fmt.Println(number_cap2)

	p := methods.Personaje{Name: "pepe", HP: 6.7, Gender: methods.Human}

	p.Describir()

	p.Curar(60)

	p.Describir()

	pointer := &p.Gender

	fmt.Println("esto es la dirección de genero del personaje 'p'", pointer)

	*pointer = "pepe"

	fmt.Println("modificación del puntero ", p.Gender)

	/*
		Hacer un cambio en el puntero no cambia la dirección de memoria, cambia el objeto almacenado
		en el espacio de memoria al cual hace referencia
	*/

}
