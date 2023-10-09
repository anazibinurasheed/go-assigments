package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of 2D array")
	fmt.Scan(&size)
	fmt.Println("Enter the values of array1")
	array1 := createArray(size)
	fmt.Println("Enter the values of array2")
	array2 := createArray(size)
	fmt.Println("Sum of 2 array is :")
	sum2dArrayandPrint(array1, array2)
}
func sum2dArrayandPrint(array1 [][]int, array2 [][]int) {
	sum := make([][]int, len(array1))
	for i := 0; i < len(array1); i++ {
		sum[i] = make([]int, len(array1))
		for j := 0; j < len(array1); j++ {
			sum[i][j] = array1[i][j] + array2[i][j]
			fmt.Print(sum[i][j], " ")
		}
		fmt.Println()
	}
}
func createArray(size int) [][]int {
	array := make([][]int, size)
	for i := 0; i < size; i++ {
		array[i] = make([]int, size)
		for j := 0; j < size; j++ {
			var input int
			fmt.Scan(&input)
			array[i][j] = input
		}
	}
	return array
}
