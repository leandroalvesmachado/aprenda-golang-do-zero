package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// isso faz com que o canal1 precise esperar o canal2 para exibir a mensagem
		// pois o tempo do canal1 é menor, estamos perdendo tempo ao esperar o canal2

		// Antes do select esta imprimindo assim
		// Canal 1
		// Canal 2
		// Canal 1
		// Canal 2
		// Canal 1
		// Canal 2
		// Canal 1
		// Canal 2
		// Canal 1

		// com o uso do select ficou assim
		// Canal 1
		// Canal 1
		// Canal 1
		// Canal 2
		// Canal 1
		// Canal 1
		// Canal 1
		// Canal 1
		// Canal 2
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
