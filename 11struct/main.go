package main

import "fmt"

func main() {
	u1 := User{"anas", 20, true}
	fmt.Println(u1)
	fmt.Printf("%+v", u1)
	fmt.Println("")
	fmt.Printf("%v %v", u1.Name, u1.Age)

}

type User struct {
	Name   string
	Age    int
	Status bool
}
