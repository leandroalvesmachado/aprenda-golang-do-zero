package main

import "fmt"

func main() {
	// ARITMETICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Go nao permite fazer operacoes com numeros de tipos diferentes
	// fortemente tipado
	// var numero1 int16 = 10
	// var numero2 int32 = 20
	// resultado := numero1 + numero2
	// fmt.Println(resultado)

	// ATRIBUICAO
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS (3)
	// && (and), || (or), !(contrario)
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!false)

	// OPERADORES UNARIOS
	numero1 := 10
	numero1++
	fmt.Println(numero1)

	numero2 := 10
	numero2 += 15
	fmt.Println(numero2)

	numero3 := 10
	numero3--
	fmt.Println(numero3)

	numero4 := 10
	numero4 -= 15
	fmt.Println(numero4)

	numero5 := 10
	numero5 *= 2
	fmt.Println(numero5)

	numero6 := 10
	numero6 /= 2
	fmt.Println(numero6)

	numero7 := 10
	numero7 %= 2
	fmt.Println(numero7)

	// OPERADOR TERNARIO NAO EXISTE
}
