package main

import (
	"fmt"
)

func main() {

	temp := hex{40, 88}
	fmt.Println(temp)
	fmt.Println(temp.area())
}

type hex struct {
	num int
	age int
}

func (args *hex) area() int {
	return args.age + args.num
}
