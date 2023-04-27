package main

import "fmt"

func main() {
	anas := User{"anas", 18, true}
	anas.getInfo()
	anas.changeName("anaz ibinu rasheed")
	anas.getInfo()
}

type User struct {
	Name   string
	Age    int
	status bool
}

func (u User) getInfo() {
	fmt.Printf("%+v", u)
}

func (u *User) changeName(n string) {
	u.Name = n

}
