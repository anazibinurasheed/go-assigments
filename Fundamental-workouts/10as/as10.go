package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 [5]int
	fmt.Println("Enter the values of array1")
	add(arr1[:])
	fmt.Println("Enter the values of array2")
	add(arr2[:])

	temp1, temp2 := swap(arr1[:], arr2[:])
	arr1 = [5]int(temp1)
	arr2 = [5]int(temp2)
	fmt.Println("After swap alues of array1")
	print(arr1[:])
	fmt.Println("After swap values of array2")
	print(arr2[:])

}
func add(arr []int) {
	for i := 0; i < len(arr); i++ {
		var input int
		fmt.Scan(&input)
		arr[i] = input

	}
}
func swap(arr1 []int, arr2 []int) ([]int, []int) {
	for i := 0; i < len(arr1); i++ {
		temp := arr1[i]
		arr1[i] = arr2[i]
		arr2[i] = temp
	}
	return arr1, arr2
}
func print(arr []int) {
	for i := 0; i < len(arr); i++ {

		fmt.Print(arr[i], " ")
	}
	fmt.Println("")
}
