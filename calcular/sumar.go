package calcular

import (
	"errors"
	"fmt"
)

type Type interface {
	int | int8 | int16 | int32 | uint |
		uint8 | uint16 | uint32 | float32
}

/*
Defino un tipo entero general para no tener que definir siempre las variables
dentro de las funciones
*/

func Sumar[T Type](a T, b T) T {
	result := a + b
	return result
}

func Restar[T Type](a T, b T) T {
	result := a - b
	return result
}

func Multiply[T Type](a T, b T) T {
	result := a * b
	return result
}

func corroboratedivision[T Type](a T, b T) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por 0")
	}

	result := float64(a) / float64(b)
	return result, nil
}

func Dividir[T Type](a, b T) {
	cociente, err := corroboratedivision(a, b)
	if err == nil {
		fmt.Println("División: ", cociente)
	} else {
		fmt.Println("Error:", err)
		return
	}

}
