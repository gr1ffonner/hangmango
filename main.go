package main

import (
	"fmt"
	"slices"
	"strings"
)

func hideWord(word string) string {
	return strings.Repeat(star, len(word))
}

func unhideWord(word string, hiddenWord string, letter string) string {
	for i, v := range word {
		if strings.ToLower(string(v)) == letter {
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
	secret, err := getWord()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	secret = strings.ToLower(secret)
	var userInput string

	fmt.Println("Helloooooo, here is a hangman game written in Golang for learning purposes")

	for {
		fmt.Print("Do you want to start a new game? [Y]es [N]o: ")
		fmt.Scanln(&userInput)

		userInput = strings.ToLower(userInput)

		if userInput != "y" && userInput != "n" {
			fmt.Println("Invalid input. Please type 'y' or 'n'.")
			continue
		}

		if strings.ToLower(userInput) == "n" {
			fmt.Println("Bye bye!!!")
			break
		}
		guesses := make([]string, 0)
		fmt.Println("Sooooo, let's start!!!")
		hiddenWord := hideWord(secret)
		fmt.Println(hiddenWord)
		incorrectGuesses := 0
		fmt.Println(hangman)
		for {
			letter := typeLetter()

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
				if incorrectGuesses == 6 {
					fmt.Printf("%s\n", hangmanStates[6])
					fmt.Println("Game over! You've made too many incorrect guesses.")
					fmt.Printf("The secret word is: %s\n", secret)
					break
				}
				printHangmanState(incorrectGuesses)
			}
		}
	}
}
