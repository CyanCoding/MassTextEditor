package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadWordsFile(path string) []string {
	byteData, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Invalid path!")
		return nil
	}

	text := string(byteData)
	words := strings.Fields(text)

	return words
}

func main() {
	for {
		fmt.Print("Please enter a text file path > ")
		var path string
		fmt.Scanln(&path)

		fmt.Println()
		fmt.Println("Reading file...")
		wordList := ReadWordsFile(path)

		if wordList != nil {
			fmt.Println("Finished reading file of ", len(wordList), " lines")
			break
		}
	}

}
