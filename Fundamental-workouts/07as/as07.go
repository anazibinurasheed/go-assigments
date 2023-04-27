package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var input int
	fmt.Scan(&input)

	for i := 1; i <= 10; i++ {
		mul := i * input
		fmt.Println(i, "x", input, "=", mul)
	}
}
