package main

import (
	"fmt"
	"strings"
)

var secret string = "danya"

const (
	star    string = "*"
	hangman string = `
     ________
     |/      |
     |     
     |    
     |   
     |  
     |
    _|___
  `
	err1 string = `
     ________
     |/      |
     |      (_)
     |      
     |     
     |    
     |
    _|___
  `
	err2 string = `
     ________
     |/      |
     |      (_)
     |       |
     |       |
     |      
     |
    _|___
  `

	err3 string = `
     ________
     |/      |
     |      (_)
     |      \|
     |       |
     |      
     |
    _|___
  `
	err4 string = `
     ________
     |/      |
     |      (_)
     |      \|/
     |       |
     |      
     |
    _|___
  `
	err5 string = `
     ________
     |/      |
     |      (_)
     |      \|/
     |       |
     |      / 
     |
    _|___
  `
	err6 string = `
     ________
     |/      |
     |      (_)
     |      \|/
     |       |
     |      / \
     |
    _|___
  `
)

func hideWord(word string) string {
	return strings.Repeat(star, len(word))
}

func unhideWord(word string, hiddenWord string, letter string) string {
	// Convert the letter to lowercase to handle case-insensitive matching
	letter = strings.ToLower(letter)

	// Use a flag to track if the letter was correctly guessed
	letterCorrectlyGuessed := false

	// Iterate over the original word to update the hidden word
	for i, v := range word {
		if strings.ToLower(string(v)) == letter {
			// Check if the letter was already correctly guessed
			if string(hiddenWord[i]) == letter {
				fmt.Println("You've already guessed the letter '", letter, "'.")
			} else {
				// Replace the corresponding '*' with the guessed letter
				hiddenWord = hiddenWord[:i] + letter + hiddenWord[i+1:]
				letterCorrectlyGuessed = true
			}
		}
	}

	// If the letter was not correctly guessed, print a message
	if !letterCorrectlyGuessed {
		fmt.Println("There is no letter '", letter, "' in the secret word.")
	}

	return hiddenWord
}

func printHangmanState(incorrectGuesses int) {
	switch incorrectGuesses {
	case 1:
		fmt.Println(err1)
		fmt.Printf("Incorrect Guesses: %d\n", incorrectGuesses)
	case 2:
		fmt.Println(err2)
		fmt.Printf("Incorrect Guesses: %d\n", incorrectGuesses)
	case 3:
		fmt.Println(err3)
		fmt.Printf("Incorrect Guesses: %d\n", incorrectGuesses)
	case 4:
		fmt.Println(err4)
		fmt.Printf("Incorrect Guesses: %d\n", incorrectGuesses)
	case 5:
		fmt.Println(err5)
		fmt.Printf("Incorrect Guesses: %d\n", incorrectGuesses)
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
			fmt.Println("Sooooo, let's start!!!")
			hiddenWord := hideWord(secret)
			fmt.Println(hiddenWord)
			incorrectGuesses := 0
			fmt.Println(hangman)
			for {
				letter := typeLetter()

				if strings.Contains(secret, letter) {
					hiddenWord = unhideWord(secret, hiddenWord, letter)
					fmt.Println(hiddenWord)

					// Check if the word is completely revealed
					if hiddenWord == secret {
						fmt.Println("Congratulations! You've revealed the word:", secret)
						break
					}
				} else {
					fmt.Println("There is no letter '", letter, "' in the secret word.")
					incorrectGuesses++

					// Print hangman state and check if the player has too many incorrect guesses
					printHangmanState(incorrectGuesses)
					if incorrectGuesses == 6 {
						fmt.Println(err6)
						fmt.Println("Game over! You've made too many incorrect guesses.")
						break
					}
				}
			}

			fmt.Print("Do you want to play again? [Y]es [N]o: ")
			fmt.Scanln(&userInput)
			if strings.ToLower(userInput) == "n" {
				fmt.Println("Bye bye!!!")
				break
			}
		}
	}
}
