package main

import (
	"fmt"
	"math/rand"
)

func main() {

	dice := rand.Intn(6) + 1

	switch dice {
	case 1:
		fmt.Println("Now you can move 1 step forward ")
	case 2:
		fmt.Println("you can move 2 steps forward")
	case 3:
		fmt.Println("you can move 3 steps forward")
	case 4:
		fmt.Println("you can move 4 steps forward")
	case 5:
		fmt.Println("you can move 5 steps forward ")
	case 6:
		fmt.Println("Now you can realease 1 dice and  you can move 6 steps forward")
	default:
		fmt.Println("What was that")
	}

}
