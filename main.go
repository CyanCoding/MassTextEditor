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

	fmt.Println("Available options:")
	fmt.Println("(1) Remove all lines with capital letters")
	fmt.Println("(2) Remove all lines with lowercase letters")
	fmt.Println("(3) Remove all lines with a certain character")
	fmt.Println("(4) Remove all lines without a certain character")
	fmt.Println("------------------------------------------------")
	for {
		fmt.Print("Action number > ")
		var action int = 0
		fmt.Scanln(&action)

		if action > 0 && action < 5 { // Action is valid
			break
		} else {
			fmt.Println()
			fmt.Println("Invalid action!")
			fmt.Println()
		}
	}

}
