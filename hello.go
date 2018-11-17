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
	wordCounter := map[string]int32{}    //create a new dictionary to insert words with its counter
	for _, element := range words {      //iterate through dictionary
		processedString := reg.ReplaceAllString(element, "") //remove special characters from string
		if len(processedString) > 0 {                        //check if the processed word is empty
			word := strings.ToLower(processedString) //convert word to lowercase
			value, ok := wordCounter[word]           //check if the word is already in the dictionary
			if ok {
				wordCounter[word] = value + 1 //if the word is already in the dictionary, then increment counter value
			} else {
				wordCounter[word] = 1 //else, insert a new word to the dictionary with initial counter value to 1
			}
		}
	}
	fmt.Println(wordCounter)
}
