package main

import "fmt"

func funcaoClosure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := funcaoClosure()
	// foi possivel chamar dessa forma, pois a propria funcao retorna outra funcao
	funcaoNova()
}
