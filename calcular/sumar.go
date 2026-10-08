package calcular

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

func Multiplicar[T Type](a T, b T) T {
	result := a * b
	return result
}

func Dividir[T Type](a T, b T) float32 {
	result := float32(a) / float32(b)
	return result
}
