package main

import "fmt"

// defer significa adiar

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func main() {
	// neste exemplo a função 1 será executada depois da função 2
	defer funcao1()
	funcao2()
}
