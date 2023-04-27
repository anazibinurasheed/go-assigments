package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size of array")
	fmt.Scan(&size)

	fmt.Println("Enter the values of array")
	array := addvalue(&size)
	countOfEven(&array)

}
func countOfEven(array *[]int) {
	var count int
	for _, val := range *array {
		if val%2 == 1 {
			count++
		}

	}
	fmt.Println("Number of even number in the given array:", count)
}

func addvalue(size *int) []int {
	array := make([]int, *size)
	for i := 0; i < *size; i++ {
		var input int
		fmt.Scan(&input)
		array[i] = input

	}
	return array
}
