package main

import "fmt"

func main() {
	func(txt string) {
		fmt.Println(txt)
	}("Passando parâmetro para função anônima")
}