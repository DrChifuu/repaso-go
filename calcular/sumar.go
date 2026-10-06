package calcular

type Entero interface {
	int | int8 | int16 | uint | uint8 | uint16
}

func Sumar[T Entero](a T, b T) T {
	resultado := a + b
	return resultado
}
