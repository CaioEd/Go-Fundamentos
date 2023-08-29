package main

import "fmt"

func main() {
	var v1 int = 100
	var ponteiro *int 		// PONTEIRO

	ponteiro = &v1

	v1++

	fmt.Println(v1, ponteiro)	// 101 0xc00000a098 - MOSTRA O VALOR DE V1 E SEU ENDEREÇO
	fmt.Println(v1, *ponteiro)	// 101 101 - MOSTRA OS VALORES
}