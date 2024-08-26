package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type Estudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Pseudo Herança")

	pessoa1 := Pessoa{"Joao", "Pedro", 20, 178}
	// {Joao Pedro 20 178}
	fmt.Println(pessoa1)

	estudante1 := Estudante{pessoa1, "engenharia", "USP"}
	// {{Joao Pedro 20 178} engenharia USP}
	fmt.Println(estudante1)

	// Joao
	fmt.Println(estudante1.nome)
}
