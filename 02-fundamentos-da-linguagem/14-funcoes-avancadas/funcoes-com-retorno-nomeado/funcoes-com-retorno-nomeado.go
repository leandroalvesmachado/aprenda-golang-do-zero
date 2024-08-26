package main

import "fmt"

// (n1, n2 int) = parametros do tipo int
// (soma int, subtracao int) = retornos e seus tipos
// não preciso criar as variaveis soma e subtracao, ja estão criadas com seus tipos e basta retornar
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("Função com Retorno Nomeado")
	// precise ter duas variaveis, pois são 2 retornos
	resultadoSoma, resultadoSubtracao := calculosMatematicos(2, 1)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
