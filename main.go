package main

import (
	"cyancoding/go-humanize"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
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

// Checks letter by letter to see if a string is fully uppercase
// if it hits a lowercase letter it returns false and exits
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Checks letter by letter to see if a string is only alphabetic
// if it hits an invalid letter it returns false and exits
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Checks letter by letter to see if a string is fully lowercase
// if it hits an uppercase letter it returns false and exits
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func WriteToFile(wordList []string, path string) bool {
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer file.Close()
	for _, line := range wordList {
		file.WriteString(line + "\n")
	}

	return true
}

func main() {
	var wordList []string
	for {
		fmt.Print("Please enter a text file path > ")
		var path string
		fmt.Scanln(&path)

		fmt.Println()
		fmt.Println("Reading file...")
		wordList = ReadWordsFile(path)

		if wordList != nil {
			fmt.Println("Finished reading file of", humanize.Comma(int64(len(wordList))), "lines")
			break
		}
	}

	fmt.Println("Available options:")
	fmt.Println("(1) Remove all lines with capital letters")
	fmt.Println("(2) Remove all lines with lowercase letters")
	fmt.Println("(3) Remove all lines with a certain character")
	fmt.Println("(4) Remove all lines without a certain character")
	fmt.Println("(5) Remove all lines with non-letters")
	fmt.Println("------------------------------------------------")
	var action int = 0
	for {
		fmt.Print("Action number > ")
		fmt.Scanln(&action)

		if action > 0 && action < 6 { // Action is valid
			break
		} else {
			fmt.Println()
			fmt.Println("Invalid action!")
			fmt.Println()
		}
	}

	var specialCharacter string
	if action == 3 {
		fmt.Println()
		fmt.Print("Please enter the character you want the file to include > ")
		fmt.Scanln(&specialCharacter)
	} else if action == 4 {
		fmt.Println()
		fmt.Print("Please enter the character you want the file to not include > ")
		fmt.Scanln(&specialCharacter)
	}

	newList := make([]string, 0)

	for i := 0; i < len(wordList); i++ {
		addWord := false

		if action == 1 && IsLower(wordList[i]) {
			addWord = true
		} else if action == 2 && IsUpper(wordList[i]) {
			addWord = true
		} else if action == 3 && strings.Contains(wordList[i], specialCharacter) {
			addWord = true
		} else if action == 4 && !strings.Contains(wordList[i], specialCharacter) {
			addWord = true
		} else if action == 5 && IsLetter(wordList[i]) {
			addWord = true
		}

		if addWord {
			newList = append(newList, wordList[i])
		}
	}

	fmt.Println()
	fmt.Println("Finished the job! Removed",
		humanize.Comma(int64(len(wordList)-len(newList))),
		"lines")

	for {
		fmt.Print("Output file path > ")
		var outputPath string
		fmt.Scanln(&outputPath)

		success := WriteToFile(newList, outputPath)

		fmt.Println()
		if success {
			fmt.Println("Thanks for using the MassTextEditor!")
		}
		if !success {
			fmt.Println("Failed to write to the file! Please try again.")
			fmt.Println()
		}
	}

}
