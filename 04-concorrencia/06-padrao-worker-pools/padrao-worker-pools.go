package main

import "fmt"

// utilizado para filas de tarefas grandes

func main() {
	// criando dois canais com buffers de 45 posições
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		// adiciona o valor a variável tarefas
		tarefas <- i
	}

	// fechando o canal tarefas
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// tarefas <-chan int (canal que recebe dados, <-chan)
// resultados chan<- int (canal que envia dados, chan<-)
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
