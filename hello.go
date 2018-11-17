package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(b)             // convert file content into a 'string'
	words := strings.Fields(fileContent) // insert each word into an array TODO: ignore special characters and process lower case
	for _, element := range words {      //iterate through dictionary
		processedString := reg.ReplaceAllString(element, "") //remove special characters from string
		if len(processedString) > 0 {                        //check if the word is empty
			fmt.Println(processedString)
		}
	}
}
