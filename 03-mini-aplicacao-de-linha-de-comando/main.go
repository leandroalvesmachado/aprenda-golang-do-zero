package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	// erro := aplicacao.Run(os.Args)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }

	// outra forma de fazer esse if
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
