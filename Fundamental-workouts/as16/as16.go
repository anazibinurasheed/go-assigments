package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	var isprime = false
	for i := 2; i < num; i++ {

		if num%i == 0 {
			isprime = true
		}
	}
	if isprime == true {
		fmt.Println("Not prime")
	} else {
		fmt.Println("Prime")
	}

}
