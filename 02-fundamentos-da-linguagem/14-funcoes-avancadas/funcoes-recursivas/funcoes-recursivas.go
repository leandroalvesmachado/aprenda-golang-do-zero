package main

import "fmt"

// uint só aceita números positivos
// int aceita positivos e negativos
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 5 8 13
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	// 144

	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
	// 89
}
