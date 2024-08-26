package main

import "fmt"

// essa função pode receber N parâmetros
func soma(numeros ...int) int {
	fmt.Println(numeros) // [1 2 3 4]
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

// combinando tipos de parâmetros
// só é possível ter 1 parâmetro variatico por função
//
//	ele precisa ser sempre o último parâmetro
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções Variáticas")
	totalSoma := soma(1, 2, 3, 4)
	fmt.Println(totalSoma) // 10

	escrever("Olá", 1, 2, 3, 4)
}
