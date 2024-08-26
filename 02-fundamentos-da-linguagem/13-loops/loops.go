package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0

	// parecido com o while
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second) // sleep de 1 segundo
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j:", j)
		time.Sleep(time.Second) // sleep de 1 segundo
	}

	vogais := [3]string{"A", "B", "C"}

	for indice, vogal := range vogais {
		fmt.Println(indice, vogal)
	}

	for _, vogal := range vogais {
		fmt.Println(vogal)
	}

	for _, letra := range "PALAVRAS" {
		fmt.Println(letra)
		// retorna a tabela ASCII da letra
		// 80 = P
		// 65 = A
		// 76
		// 65
		// 86
		// 82
		// 65
		// 83

		// se quiser a letra mesmo
		fmt.Println(string(letra)) // P, A, L, ...
	}

	// estrutura map
	usuario := map[string]string{
		"nome":      "A",
		"sobrenome": "B",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// nao é possivel utiliza o range no struct

	// loop infinito
	for true {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
