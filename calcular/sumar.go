package calcular

type Entero interface {
	int | int8 | int16 | int32 | uint |
		uint8 | uint16 | uint32 | float32
}

/*
Defino un tipo entero general para no tener que definir siempre las variables
dentro de las funciones
*/

func Sumar[T Entero](a T, b T) T {
	resultado := a + b
	return resultado
}

func Restar[T Entero](a T, b T) T {
	resultado := a + b
	return resultado
}

func Multiplicar[T Entero](a T, b T) T {
	resultado := a * b
	return resultado
}

func Dividir[T Entero](a T, b T) float32 {
	resultado := float32(a) / float32(b)
	return resultado
}
