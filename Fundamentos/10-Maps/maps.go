package main

import "fmt"

func main() {
	usuario := map[string]string {
		"nome": "Alexandre",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "Lionel",
			"Sobrenome": "Messi",
		},
	}

	fmt.Println(usuario2)
}