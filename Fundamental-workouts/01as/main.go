package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a  charecter ")
	input, _ := reader.ReadString('\n')
	fmt.Println("output:\n", input)
}
