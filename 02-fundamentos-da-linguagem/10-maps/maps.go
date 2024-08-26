package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// definindo o tipo das chaves e dos valores no map
	// nao tem como alterar os tipos do map, imutavel
	usuario1 := map[string]string{
		"nome":      "Joao",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario1)         // map[nome:Joao sobrenome:Alves]
	fmt.Println(usuario1["nome"]) // Joao

	usuario2 := map[int]string{
		1: "Joao",
		2: "Alves",
	}

	fmt.Println(usuario2)    // map[1:Joao 2:Alves]
	fmt.Println(usuario2[1]) // Joao

	// deletar chave do map usuario2 da chave 1
	delete(usuario2, 1)
	fmt.Println(usuario2) // map[2:Alves]

	usuario3 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Leandro",
		},
	}

	fmt.Println(usuario3) // map[nome:map[primeiro:Joao ultimo:Leandro]]
}
