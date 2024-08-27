package main

import (
	"fmt"
	"time"
)

func main() {
	// informando ao canal que só pode trafegar dados do tipo string
	canal := make(chan string)

	go escrever("Olá mundo", canal)
	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	// isso aqui bloqueia a execução do código, enquanto a var mensagem não receber o valor do canal, mesmo com a goroutine
	// 	mensagem, aberto := <-canal // informando que o canal esta esperando receber o valor e depois vai salvar na variavel mensagem

	// 	// canal esta aberto?
	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	// faz a mesma coisa do código acima
	// ja verifica se o canal esta aberto ou não com o range
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto // como enviar dado para o canal (mandando um valor para o canal)
		time.Sleep(time.Second)
	}

	close(canal) // fechando o canal
}
