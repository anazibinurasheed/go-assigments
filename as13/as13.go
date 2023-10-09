package main

import "fmt"

func main() {
	var input string
	var copy string
	fmt.Println("Enter a word")
	fmt.Scan(&input)
	copy = input
	if copy == input {
		fmt.Println("Entered string is palindrome")
	} else {
		fmt.Println("Entered string is not a palindrome")
	}

}
