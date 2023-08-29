package main

import "fmt"

func main() {
	idade := 10

	if idade >= 18 && idade < 60{
		fmt.Println("Maior de idade!")
	} else if idade >= 60 {
		fmt.Println("Idoso")	
	} else {
		fmt.Println("Menor de idade!")
	}

	if numero := idade; numero > 0 {
		fmt.Println("Número é maior que 0")
	} else {
		fmt.Println("ERRO Idade tem que ser maior que 0!")
	}
}