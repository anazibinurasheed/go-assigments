package main

import "fmt"

func main() {
	fmt.Println("hello world ")

	result := proAdder(88, 99, 33, 0, 1, 2)
	fmt.Println(result)

	a, b := adder(70, 88)
	fmt.Println(a, b)

}
func proAdder(values ...int) int {
	total := 0
	fmt.Printf("%T", values)
	fmt.Println("")
	for _, v := range values {
		total += v

	}
	return total
}

func third() {
	fmt.Println("third")
}

func sec() {

	fmt.Println("hi")
	third()
}
func adder(a int, b int) (int, string) {
	return a + b, "yeah its done"
}
