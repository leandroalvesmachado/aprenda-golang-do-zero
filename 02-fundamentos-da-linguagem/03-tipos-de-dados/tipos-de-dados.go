package main

import (
	"errors"
	"fmt"
)

func main() {
	// ### INTEIROS ###

	// int - define o int de acordo com a arquitetura do computador
	// int8 - is the set of all signed 8-bit integers. Range: -128 through 127
	// int16 - is the set of all signed 16-bit integers. Range: -32768 through 32767
	// int32 - is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647
	// int64 - is the set of all signed 64-bit integers. Range: -9223372036854775808 through 9223372036854775807.

	var numero int8 = 100
	fmt.Println(numero)

	// uint - define o int sem ignorando o sinal negativo
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias

	// int32 = rune
	var numero3 int32 = 123456
	var numero4 rune = 123456
	fmt.Println(numero3, numero4)

	// byte = uint8
	var numero5 byte = 123
	fmt.Println(numero5)

	// ### FLOAT ###

	// float32 - is the set of all IEEE 754 32-bit floating-point numbers.
	// float64 - is the set of all IEEE 754 64-bit floating-point numbers.
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 123.456
	numeroReal3 := 123.456
	fmt.Println(numeroReal1, numeroReal2, numeroReal3)

	// ### STRINGS ###

	var str string = "Texto 1"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// se declarar com aspa simples, ele vai retornar o codigo da tabela ASCII
	// B (char) = 42 (hex) = 66 (decimal)
	char := 'B'
	fmt.Println(char)

	// retorna uma string em branco
	var texto string
	fmt.Println(texto)

	// retorna 0
	var inteiro int16
	fmt.Println(inteiro)

	// ### BOOLEANOS ###

	// valor padrao = false
	var booleano bool
	fmt.Println(booleano)

	// nil
	var erro1 error
	fmt.Println(erro1)

	// Erro ao cadastrar o produto
	var erro2 error = errors.New("Erro ao cadastrar o produto")
	fmt.Println(erro2)
}
