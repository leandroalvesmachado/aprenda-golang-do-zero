package main

import "fmt"

func main() {
	// formas de declarar variaveis

	// variavel do tipo string declarada
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	// variavel do tipo string declarada com o uso da inferencia de tipos
	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	// constantes nao podem ter seu valor alterado
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// trocando valor das variaveis
	variavel5, variavel6 = variavel1, variavel2
	fmt.Println(variavel5, variavel6)
}
