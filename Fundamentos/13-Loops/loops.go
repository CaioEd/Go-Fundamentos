package main

import (
	"fmt"
)

func main() {
	usuarios := [3]string{"Caie", "Davi", "Ana"}

	for indice, user := range usuarios {
		fmt.Println(indice, user)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}