package main

import "fmt"

func main() {
	slice1 := []int{90, 78, 534, 676}
	fmt.Println(slice1)

	slice1 = append(slice1, 19)
	fmt.Println(slice1)

	array1 := [5]int{1, 4, 8, 10, 14}
	slice2 := array1[1:4]
	fmt.Println(slice2)
}