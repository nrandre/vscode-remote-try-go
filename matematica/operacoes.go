package matematica

import "strconv"

func Soma(a int, b int) string {
	calculo := a + b
	resultado := strconv.Itoa(calculo)

	return resultado
}
