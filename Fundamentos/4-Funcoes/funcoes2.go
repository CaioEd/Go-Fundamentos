package main

import "fmt"

func calculos(n1, n2 int) (int, int) {	
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	resultadoSoma, resultadoSubtracao := calculos(349, 122)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}