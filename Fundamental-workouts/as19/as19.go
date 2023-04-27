package main

import "fmt"

func main() {
	var income float32
	var tax float32
	fmt.Println("Enter your income")
	fmt.Scan(&income)
	if income <= 250000 {
		fmt.Println("No tax")
	} else if income <= 500000 {
		tax = income * 0.05
	} else if income <= 1000000 {
		tax = income * 0.20
	} else if income <= 5000000 {
		tax = income * 0.30

	} else {
		fmt.Println("Invalid entry")
	}

	fmt.Printf("Income tax amount%.f:", tax)

}
