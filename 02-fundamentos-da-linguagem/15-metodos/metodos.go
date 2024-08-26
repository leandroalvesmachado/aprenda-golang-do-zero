package main

import "fmt"

func escrever() {
	fmt.Println("Escrevendo...")
}

type Usuario struct {
	nome  string
	idade uint8
}

func (u Usuario) salvar() {
	fmt.Println(u.nome, u.idade)
	fmt.Println("Dentro do método salvar")
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u Usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// utilizando os ponteiros, eu consigo alterar o valor do objeto Usuario que chamou a função no atributo idade
func (u *Usuario) fazerAniversario() {
	u.idade++
}

func main() {
	escrever()

	usuario1 := Usuario{"Pedro", 20}
	fmt.Println(usuario1)
	usuario1.salvar()                    // chamando o método salvar
	fmt.Println(usuario1.maiorDeIdade()) // chamando o método maiorDeIdade

	fmt.Println("\n")

	usuario2 := Usuario{"Davi", 17}
	fmt.Println(usuario2)
	usuario2.salvar()                    // chamando o método salvar
	fmt.Println(usuario2.maiorDeIdade()) // chamando o método maiorDeIdade

	fmt.Println("\n")

	usuario2.fazerAniversario()
	// objeto usuario2 foi realmente alterado o atributo idade
	fmt.Println(usuario2) // {Davi 18}
}
