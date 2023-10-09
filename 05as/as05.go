package main

import "fmt"

func main() {
	fmt.Println("Enter your mark")
	var totalmark float32
	fmt.Scan(&totalmark)

	if totalmark >= 90 && totalmark <= 100 {
		fmt.Println("A")
	} else if totalmark >= 80 && totalmark < 90 {
		fmt.Println("B")
	} else if totalmark >= 70 && totalmark < 80 {
		fmt.Println("C")
	} else if totalmark >= 60 && totalmark < 70 {
		fmt.Println("D")
	} else if totalmark >= 50 && totalmark < 60 {
		fmt.Println("E")
	} else if totalmark >= 0 && totalmark < 50 {
		fmt.Println("FAILED")
	} else {
		fmt.Println("You have entered a invalid mark")
	}

}
