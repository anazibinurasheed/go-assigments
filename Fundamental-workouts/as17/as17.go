package main

import "fmt"

func main() {
	fmt.Println("Enter 2 ")
	worker()
}

type math struct {
	num1 float32
	num2 float32
}

func (val math) addition() float32 {

	return val.num1 + val.num2
}
func (val math) multiplication() float32 {

	return val.num1 * val.num2
}
func (val math) division() float32 {

	return val.num1 / val.num2
}

func (val math) subtraction() float32 {

	return val.num1 - val.num2
}
func worker() {
	var choice int
	var number1 float32
	var number2 float32
	fmt.Println("Select your operation \n 1:Addition \n 2:Multiplication \n 3:Subtraction \n 4:Division")
	fmt.Scan(&choice)
	fmt.Println("Enter 2 numbers")
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	userinput := math{number1, number2}

	switch choice {
	case 1:
		fmt.Println("Result:", userinput.addition())
	case 2:
		fmt.Println("Result:", userinput.multiplication())
	case 3:
		fmt.Println("Result:", userinput.subtraction())
	case 4:
		fmt.Println("Result:", userinput.division())
	default:
		fmt.Println("Invalid entry!!!")

	}

}
