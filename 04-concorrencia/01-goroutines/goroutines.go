package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO (tarefas ao mesmo tempo)

	// goroutine, executa essa função escrever(), mas não precisa esperar para executar as próximas instruções
	go escrever("Olá mundo")
	escrever("Programando em GO")

	// exemplo de saída
	// Programando em GO
	// Olá mundo
	// Olá mundo
	// Programando em GO
	// Olá mundo
	// Programando em GO
	// Olá mundo
	// Programando em GO
	// Olá mundo
	// Programando em GO
}

func escrever(texto string) {
	// for infinito
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
