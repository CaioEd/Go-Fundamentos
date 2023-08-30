package main

import "fmt"

type user struct {
	nome string
	idade uint
}

func (u user) salvar() {
	fmt.Printf("Salvando dados do usuário %s!", u.nome)
}

func (u *user) aniversario() {
	u.idade++
}

func main() {
	user1 := user{"Dean", 28}
	fmt.Println(user1)
	user1.salvar()

	user2 := user{"Sam", 23}
	fmt.Println(user2)
	
	user2.aniversario()
	fmt.Println(user2)
}