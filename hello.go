package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("hello world")
	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(b)             // convert file content into a 'string'
	words := strings.Fields(fileContent) // insert each word into an array TODO: ignore special characters and process lower case
	for _, element := range words {      //iterate through dictionary
		fmt.Println(element)
	}
}
