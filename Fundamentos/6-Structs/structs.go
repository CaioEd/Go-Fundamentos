package main

import "fmt"

type user struct {
	nome string
	idade int
}

func main() {
	user1 := user{"David", 25}
	fmt.Println(user1)

	user2 := user{nome: "Chris"}
	fmt.Println(user2)
}