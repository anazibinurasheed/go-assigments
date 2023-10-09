package main

import "fmt"

func main() {
	array := declareArray()
	getArray(&array)
	printArray(array)

}
func declareArray() []int {
	var array []int
	return array
}
func getArray(array *[]int) {
	*array = append(*array, 12, 30, 40, 50, 60)

}
func printArray(array []int) {
	for _, val := range array {
		fmt.Print(val, " ")
	}
}
