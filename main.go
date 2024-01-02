package main

import (
	"fmt"
	"slices"
	"strings"
)

var (
	secret string = "danya"

	hangmanStates = map[int]string{
		1: `
     ________
     |/      |
     |      (_)
     |
     |
     |
     |
    _|___
  `,
		2: `
     ________
     |/      |
     |      (_)
     |       |
     |       |
     |
     |
    _|___
  `,
		3: `
     ________
     |/      |
     |      (_)
     |      \|
     |       |
     |
     |
    _|___
  `,
		4: `
     ________
     |/      |
     |      (_)
     |      \|/
     |       |
     |
     |
    _|___
  `,
		5: `
     ________
     |/      |
     |      (_)
     |      \|/
     |       |
     |        \
     |
    _|___
  `,
		6: `
     ________
     |/      |
     |      (_)
     |      \|/
     |       |
     |      / \
     |
    _|___
  `,
	}
)

const (
	alphabet string = "abcdefghijklmnopqrstuvwxyz"
	star     string = "*"
	hangman  string = `
     ________
     |/      |
     |     
     |    
     |   
     |  
     |
    _|___
  `
)

func hideWord(word string) string {
	return strings.Repeat(star, len(word))
}

func unhideWord(word string, hiddenWord string, letter string) string {
	// Iterate over the original word to update the hidden word
	for i, v := range word {
		if strings.ToLower(string(v)) == letter {
			// Check if the letter was already correctly guessed
			hiddenWord = hiddenWord[:i] + letter + hiddenWord[i+1:]
		}
	}

	return hiddenWord
}

func printHangmanState(incorrectGuesses int) {
	if state, exists := hangmanStates[incorrectGuesses]; exists {
		fmt.Printf("%s\nIncorrect Guesses: %d\n", state, incorrectGuesses)
	}
}

func typeLetter() string {
	var letter string
	fmt.Print("Type letter: ")
	fmt.Scanln(&letter)
	return strings.ToLower(letter)
}

func main() {
	var userInput string

	fmt.Println("Helloooooo, here is a hangman game written in Golang for learning purposes")

	for {
		fmt.Print("Do you want to start a new game? [Y]es [N]o: ")
		fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "n" {
			fmt.Println("Bye bye!!!")
			break
		} else {
			guesses := make([]string, 0)
			fmt.Println("Sooooo, let's start!!!")
			hiddenWord := hideWord(secret)
			fmt.Println(hiddenWord)
			incorrectGuesses := 0
			fmt.Println(hangman)
			for {
				letter := typeLetter()
				letter = strings.ToLower(letter)

				if !strings.Contains(alphabet, letter) {
					fmt.Println("Invalid letter. Please try again.")
					continue
				}

				if slices.Contains(guesses, letter) {
					fmt.Println("You've already guessed this letter. Try again.")
					continue
				}

				if strings.Contains(secret, letter) {
					hiddenWord = unhideWord(secret, hiddenWord, letter)
					guesses = append(guesses, letter)
					fmt.Println("Your guesses: ", guesses)
					fmt.Println(hiddenWord)

					// Check if the word is completely revealed
					if hiddenWord == secret {
						fmt.Println("Congratulations! You've revealed the word:", secret)
						break
					}
				} else {
					fmt.Println(hiddenWord)
					fmt.Println("There is no letter '", letter, "' in the secret word.")
					guesses = append(guesses, letter)
					fmt.Println("Your guesses: ", guesses)
					incorrectGuesses++
					printHangmanState(incorrectGuesses)
					if incorrectGuesses == 6 {
						fmt.Printf("%s", hangmanStates[6])
						fmt.Println("Game over! You've made too many incorrect guesses.")
						break
					}
				}
			}
		}
	}
}
