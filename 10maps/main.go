package main

import (
	"fmt"
)

func main() {
	table := make(map[string]int)
	table["anas"] = 8590138151
	table["muni"] = 8594126588
	fmt.Println(table)
	fmt.Println(table["anas"])
	delete(table, "anas")
	fmt.Println(table["nas"])

	for key, value := range table {
		fmt.Printf("key %v , value %v", key, value)
	}

}
