package main

import "fmt"

func soma(numeros ...int)  {
	fmt.Println(numeros)
}

func main() {
	soma(1, 20, 400, 77)
}