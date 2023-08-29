package main

import "fmt"

func main() {
	var array1[5] int 
	fmt.Println(array1)

	array2 := [3]int{10, 20, 77}
	fmt.Println(array2)

	array3 := [...]int{10, 66, 77, 999}
	fmt.Println(array3)
}