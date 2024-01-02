package main

import (
	"bufio"
	"math/rand"
	"os"
)

const filePath string = "words.txt"

func getWord() (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read lines from the file
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Select a random word
	randomIndex := rand.Intn(len(words))
	selectedWord := words[randomIndex]

	return selectedWord, nil
}
