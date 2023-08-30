package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1{
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	// 1 1 2 3 5 8 13

	posicao := uint(12)

	fmt.Println(fibonacci(posicao))
}