package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {

	result, err := div(1000, 0)

	if err != nil {
		fmt.Println("Err")
		fmt.Println(err)
		os.Exit(-1)

	} else {
		fmt.Println(result)
	}

	a := time.Now().Format("01-02-2006  Monday 15:04:05")
	createdDate := time.Date(2003, time.August, 7, 13, 1, 0, 0, time.UTC)
	fmt.Println(a)
	fmt.Println(createdDate)
}
func div(x, y int) (float64, error) {
	if y == 0 {
		return float64(0), errors.New("Cant be divided with the given input ")
	} else {
		return float64(x / y), nil
	}
}
