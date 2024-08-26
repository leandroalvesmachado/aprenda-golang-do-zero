package main

import "fmt"

type Usuario struct {
	nome     string
	idade    int8
	endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u Usuario
	u.nome = "Leandro"
	u.idade = 30

	enderecoExemplo := Endereco{"Rua teste", 123}

	// {Leandro 30}
	fmt.Println(u)

	usuario2 := Usuario{nome: "Leandro", idade: 30}
	fmt.Println(usuario2)

	usuario3 := Usuario{"Leandro", 30, enderecoExemplo}
	// {Leandro 30 {Rua teste 123}}
	fmt.Println(usuario3)

	usuario4 := Usuario{idade: 30}
	// { 30}
	fmt.Println(usuario4)
}
