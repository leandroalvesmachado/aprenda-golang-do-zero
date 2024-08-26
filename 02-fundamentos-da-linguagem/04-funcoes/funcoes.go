package main

import "fmt"

// criando uma funcao com parametros e especificando o seu retorno como int8
func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// informando que os parametros tem o mesmo tipo int8
// estou informando que a funcao vai ter dois retornos e com seus tipos int8 e int8 respectivamente
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := soma(10, 20)
	fmt.Println(soma)

	// criando funcao em uma variavel
	var f = func() {
		fmt.Println("função f")
	}

	// chamando funcao f
	f()

	// resultado := f()
	// fmt.Println(resultado)

	// funcao que tem dois retornos, entao precisa de duas variaveis para receber o retorno
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// usar underline para ignorar o segundo retorno
	// pois o go nao deixa criar variaveis e nao utilizar
	resultadoSomaComUmRetorno, _ := calculosMatematicos(5, 30)
	fmt.Println(resultadoSomaComUmRetorno)
}
