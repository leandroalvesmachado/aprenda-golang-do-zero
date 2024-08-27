package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// tem 2 goroutines para executar
	// serve para sincronizar as goroutines, nesse caso, 2 goroutines
	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em GO")
		waitGroup.Done() // -1
	}()

	// para esperar a contagem das goroutines chegar em zero (0)
	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
