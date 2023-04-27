package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your pizza rating")
	input, _ := reader.ReadString('\n')

	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("err")

	} else {
		fmt.Println("we added +1 to the rating ", value+1)

	}
}
