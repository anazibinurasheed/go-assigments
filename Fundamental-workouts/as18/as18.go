package main

import "fmt"

func main() {
	var grade float32
	var wt int
	var lb int
	var as int
	fmt.Println("Enter the mark below \n \n  ")
	fmt.Println("Written test")
	fmt.Scan(&wt)
	fmt.Println("Lab exams ")
	fmt.Scan(&lb)
	fmt.Println("Assignments")
	fmt.Scan(&as)

	grade = (float32(wt)*70)/100 + (float32(lb)*20)/100 + (float32(as)*10)/100
	fmt.Println("Grade of the student is:", grade)

}
