package main

import "fmt"

func somar(n1 int, n2 int) int {	// INT APÓS OS PARENTESES INDICA O TIPO DO RETORNO DA FUNÇÃO
	return n1 + n2
}

func main() {
	soma := somar(77, 127)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função dentro da variável")
	}

	f()
}