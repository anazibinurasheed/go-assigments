package main

import "fmt"

// 2. Accept two inputs from the user and output their sum.

func main() {
	var num1 int32
	var num2, sum float32

	fmt.Println("Enter 2 numbers : Enter the second number after a space")
	fmt.Scanf("%d%f", &num1, &num2)
	sum = add(float32(num1), num2)
	fmt.Println(sum)
}
func add(num1 float32, num2 float32) float32 {
	return num1 + num2
}
