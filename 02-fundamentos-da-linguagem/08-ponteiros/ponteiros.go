package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	// 10 10
	fmt.Println(variavel1, variavel2)

	variavel1++

	// 11 10
	// mudou apenas a 1, pois as variaveis possuem endereco de memoria diferente
	// copy
	fmt.Println(variavel1, variavel2)

	// Ponteiro = referencia de memoria
	var variavel3 int = 100
	var ponteiro *int // criando uma variavel com ponteiro

	// 100 <nil>
	fmt.Println(variavel3, ponteiro)

	variavel3++
	ponteiro = &variavel3 // jogando o endereco de memoria da varialvel3 na variavel ponteiro

	// 101, 0xc0000121d0 (endereco de memoria da variavel3)
	fmt.Println(variavel3, ponteiro)
}
