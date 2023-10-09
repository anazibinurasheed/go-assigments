/*
	Write a program to find the simple interest.

Program should accept 3 inputs from the user and calculate simple interest for the given inputs. Formula: SI=(P*R*n)/100)
Variable
Data Type
Principal amount (P)
Integer
Interest rate (R)
Float
Number of years (n)
Float
Simple Interest (SI)
Float
/
*/
package main

import (
	"fmt"
)

func main() {
	var P int
	var R, n, SI float32

	fmt.Println("Enter the principle amount")
	fmt.Scan(&P)

	fmt.Println("Enter the interest rate")

	fmt.Scan(&R)
	fmt.Println("Enter the number of years")
	fmt.Scan(&n)
	SI = ((float32(P) * R * n) / 100)
	fmt.Printf("%f", SI)
}
