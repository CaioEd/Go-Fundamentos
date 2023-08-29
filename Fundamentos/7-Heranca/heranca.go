package main

import "fmt"

type user struct {
	nome string
	idade int
	sexo string
}

type estudante struct{
	user
	curso string
	faculdade string
}

func main() {
	user1 := user{"Breno", 19, "Masculino"}

	estudante1 := estudante{user1, "ADS", "Impacta"}

	fmt.Println(estudante1)
}