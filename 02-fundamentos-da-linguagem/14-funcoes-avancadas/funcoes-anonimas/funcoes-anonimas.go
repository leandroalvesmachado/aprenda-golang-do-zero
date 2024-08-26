package main

import "fmt"

func main() {
	fmt.Println("Funções Anônimas")

	// função anônimas, declarada e executando
	func() {
		fmt.Println("Olá Mundo")
	}()

	// com parâmetro
	func(texto string) {
		fmt.Println(texto)
	}("Parâmetro 1")

	// com parâmetro e retorno
	// %s para string no Sprintf
	// %d para número no Sprintf
	resultado := func(texto string) string {
		return fmt.Sprintf("Recebido: %s", texto)
	}("Texto")

	fmt.Println(resultado)
}
