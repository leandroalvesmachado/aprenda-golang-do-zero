package main

import "fmt"

// Executando a função init
// Função main sendo executada
// Função init executa antes da main
// Cada arquivo pode conter uma função init

func init() {
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
}
