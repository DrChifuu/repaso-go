package methods

import "fmt"

type Genero string

const (
	Human  Genero = "humano"
	Goblin Genero = "goblin"
)

type Personaje struct {
	Name   string
	HP     float32
	Gender Genero
}

func (p Personaje) Describir() string {
	describe := fmt.Sprintf("\n%s tiene %.1f HP\nraza: %s", p.Name, p.HP, p.Gender)
	fmt.Println(describe)
	return describe
}

/*
	En la función usamos Sprintf y no print porque print solo imprime texto
	y en la función le prometí devolverme un string lo cual Sprintf logra
	transformando los valores a formato string y devolviendolos a la función
*/

func (p *Personaje) Curar(cantidad float32) {
	p.HP += cantidad
	fmt.Println()
	fmt.Println(cantidad, "de HP agregado")
}

/*
	Podriamos entrar en situaciones más profundas de limites de vida o vida maxima
	y vida actual para aplicar condicionales de no otorgar más vida que la máxima posible
	pero no es necesario para este caso por lo cual quedará en el simplismo
*/
