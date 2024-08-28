package main

import "fmt"

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	// canal := make(chan string)
	// canal <- "Olá Mundo"
	// mensagem := <-canal
	// fmt.Println(mensagem)

	// canal com buffer
	// capacidade do buffer igual a 2
	// assim funciona
	canal := make(chan string, 2)
	canal <- "Olá Mundo 1"
	canal <- "Olá Mundo 2"
	// canal <- "Olá Mundo 3" // erro, pois o buffer é de tamanho 2
	mensagem := <-canal
	fmt.Println(mensagem)
}
