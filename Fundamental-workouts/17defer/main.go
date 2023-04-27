package main

import (
	"fmt"
	"math/rand"
)

func main() {
	defer fmt.Println("Program is finished")

	fmt.Println("Hello world")

	for i := 0; i < 100; i++ {
		defer fmt.Printf("%d", i)
		V := rand.Intn(6) + 1

		if i == 0 {
			fmt.Println("Started the loop")
		}
		fmt.Println(V)
	}

}
