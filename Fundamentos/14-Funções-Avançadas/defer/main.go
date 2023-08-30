package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a primeira função!")
}

func funcao2() {
	fmt.Println("Executando a segunda função!")
}

func main() {
	defer funcao1()
	funcao2()
}