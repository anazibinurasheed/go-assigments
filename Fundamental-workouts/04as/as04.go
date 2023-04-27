package main

import "fmt"

func main() {
	fmt.Println("Enter your mark")
	var mark float32
	fmt.Scan(&mark)

	if mark < 50 && mark > 0 {
		fmt.Println("Failed")
	} else if mark >= 50 && mark <= 100 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Enter a valid mark")
	}

}
