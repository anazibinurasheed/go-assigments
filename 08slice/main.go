package main

import (
	"fmt"
	"sort"
)

func main() {

	var list = []string{"apple", "carrot"}
	fmt.Println(list)

	list1 := append(list)
	list1 = append(list1, list...)

	fmt.Println(list1)
	fmt.Printf("%T", list1)

	var list3 = append(list1[1:3])
	fmt.Println(list3)

	highScores := make([]int, 1)
	highScores[0] = 10
	highScores = append(highScores, 88, 889, 77)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	highScores = append(highScores, highScores...)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(highScores)

	var slice []int
	slice = append(slice, 88)
	fmt.Println("this is slice0", slice)

	var courses = []string{"reactjs", "javascript", "ruby", "golang", "swift"}
	//we want to delete ruby from the slice
	var index = 2
	fmt.Println(courses)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
