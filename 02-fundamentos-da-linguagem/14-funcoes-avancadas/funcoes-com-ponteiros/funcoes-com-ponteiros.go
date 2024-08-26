package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

// * = ponteiro (aponta sempre para o endereço de memoria)
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	// inverteu o sinal do numero, sem alterar a variavel original "numero"
	// passando um valor po copy (copia)
	numero := 20
	numeroInvertido := inverterSinal(numero) // aqui ela cria uma nova variavel
	fmt.Println(numeroInvertido)             // -20
	fmt.Println(numero)                      // 20

	// inverteu o sinal do numero, alterando tb no endereço de memoria
	// com ponteiros mexer no valor original, pois altera o valor no endereço de memoria
	novoNumero := 40
	inverterSinalComPonteiro(&novoNumero) // passando o endereço de memoria da variavel
	fmt.Println(novoNumero)               // -40
}
