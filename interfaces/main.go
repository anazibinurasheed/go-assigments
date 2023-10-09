package main

import "fmt"

func main() {
	horoshima := book{title: "hiroshima",
		price: 77,
	}

	coffee := drink{
		name:  "brucoffee",
		price: 89.89,
	}
	// a.printInfo()
	// b.printInfo()

	var info []printer
	info = append(info, horoshima, coffee)
	info[0].printInfo()
	info[1].printInfo()
	fmt.Println(info)
	rubixcube := toy{
		name:  "rubixcube",
		price: 100,
	}
	info = append(info, rubixcube)
	info[2].printInfo()
	empty(info[0])

}

type book struct {
	title string
	price float32
}
type drink struct {
	name  string
	price float32
}

func (b book) printInfo() {
	fmt.Printf("title:%s, price:%.2f", b.title, b.price)
	fmt.Println("")
}
func (c drink) printInfo() {
	fmt.Printf("name: %s, price :%f ", c.name, c.price)
	fmt.Println("")
}

type printer interface {
	printInfo()
}
type toy struct {
	name  string
	price float32
}

func (t toy) printInfo() {
	fmt.Printf("name:%s , price: %.2f", t.name, t.price)
}
func empty(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("Hi i am a string")
	case int:
		fmt.Printf("I am a integer :%d", i)
	case float32:
		fmt.Println("I am 32 bit floating point number ", i)
	case book:
		fmt.Println("i am a book")
	case drink:fmt.Println("drink")


	}
}
