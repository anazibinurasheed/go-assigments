package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This need to go  anazibinurasheed@dev.go"
	file, err := os.Create("./content.txt")
	checkNilError(err)
	io.WriteString(file, content)
	readfile(file)
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readfile(filename string) {
	ioutil.ReadFile(filename)
	fmt.Println(string(filename))
}
