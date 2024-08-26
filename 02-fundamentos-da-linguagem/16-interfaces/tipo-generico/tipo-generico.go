package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// funciona com qualquer tipo
	generica("String")
	generica(1)
	generica(true)

	// chave de qualquer tipo e valor de qualquer tipo
	mapa := map[interface{}]interface{}{
		1:   "b",
		"a": 1,
		"c": true,
	}

	fmt.Println(mapa) // map[a:1 c:true 1:b]
}
