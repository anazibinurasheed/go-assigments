package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter a number")
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println(input)

	if input < 100 {
		fmt.Println("its Ok")
	} else {
		fmt.Println("hmm")
	}

}
