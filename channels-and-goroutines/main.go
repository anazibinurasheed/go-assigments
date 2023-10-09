package main

import (
	"fmt"
	"net/http"
)

func main() {
	var links = []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://youtube.com",
	}

	c := make(chan string)
	for _, val := range links {
		go checklink(val, c)
	}

	for l:= range c{
	go	checklink(l,c)
	}
	
}
func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link)
		fmt.Println("its dowm")
		c <- link
		return

	}
	fmt.Println(link)
	fmt.Println("it up")
	c <- link
}
