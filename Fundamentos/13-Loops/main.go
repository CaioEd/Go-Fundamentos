package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println(j)
		time.Sleep(time.Second)
	}
}