package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// criando array de tamanho 5, todos os elementos devem ser do mesmo tipo
	var array1 [5]int

	// [0 0 0 0 0]
	fmt.Println(array1)

	array1[0] = 23

	// [23 0 0 0 0]
	fmt.Println(array1)

	array2 := [5]string{}

	// [    ]
	fmt.Println(array2)

	array3 := [5]string{"A", "B", "C"}

	// [A B C  ]
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5}

	// [1 2 3 4 5]
	fmt.Println(array4)

	// slice parece um array de tamanho dinamico
	slice := []int{10, 20, 30}

	// [10 20 30]
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array4)) // [5]int
	fmt.Println(reflect.TypeOf(slice))  // []int

	// funcao append = adiciona um item no final do slice
	slice = append(slice, 90)

	// [10 20 30 90]
	fmt.Println(slice)

	// capturando uma fatia do array, valores da posicao 1 ate 2
	// array3 = {"A", "B", "C"}
	slice2 := array3[1:3]
	// [B C]
	fmt.Println(slice2)

	// Arrays internos
	// []float32, 10, 15 (tipo, qtd de elementos, capacidade maxima de elementos)
	// quando a capacidade maxima é execidida, o slice aumenta a capacidade
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)      // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(slice3)) // tamanho igual a 10
	fmt.Println(cap(slice3)) // capacidade igual a 15

	slice3 = append(slice3, 20)
	fmt.Println(slice3)
}
