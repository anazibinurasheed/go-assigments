package main

import "fmt"

func main() {
	fmt.Println("Enter a the limit")
	var input int
	fmt.Scan(&input)
	sum := 0
	for i := 1; i < input; i++ {
		if i%2 == 1 {
			sum += i
			fmt.Println(sum)
		}

	}
}
