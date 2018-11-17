package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("hello world")
	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(b) // convert file content into a 'string'
	fmt.Println(fileContent)
}
