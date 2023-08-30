package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func mediaAluno(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1+n2)/2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É 6!")
}

func main() {
	fmt.Println(mediaAluno(6, 6))
	fmt.Println("Pós execução")
}