package main

import "fmt"

func main() {

	// var slice = []string{"bmw", "benz", "audi"}

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i, v := range slice {
	// 	fmt.Println(i+1, v)
	// }
	i := 0

	for i < 10 {
		if i == 5 {
			goto anas
		}

		if i == 1 {

			i++
			continue
		}
		fmt.Println(i)
		i++
	}
anas:
	fmt.Println("anazibinurasheed@gmail.com")
}
