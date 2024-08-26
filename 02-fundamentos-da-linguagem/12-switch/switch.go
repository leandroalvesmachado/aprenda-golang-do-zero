package main

import "fmt"

// não tem o break, o GO ja para automaticamente o case

func diaDaSemana1(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	default:
		return "Outro dia"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	default:
		return "Outro dia"
	}
}

func diaDaSemana3(numero int) string {
	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	default:
		diaDaSemana = "Outro dia"
	}

	return diaDaSemana
}

func diaDaSemana4(numero int) string {
	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough // o resultado desse caso passa o resultado utulizado na proxima condição
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	default:
		diaDaSemana = "Outro dia"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia1 := diaDaSemana1(200)
	fmt.Println(dia1)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(1)
	fmt.Println(dia3)
}
