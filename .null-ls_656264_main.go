package main

import (
	"fmt"
	"strings"
)

const (
	secret string = "danya"
	star   string = "*"
)

func hideWord(word string) {
	word_lenght := len(word)
	hided_word := strings.Repeat(star, word_lenght)
	fmt.Println(hided_word)
}

func typeLetter() rune {
	var letter rune
	fmt.Println("Type letter: ")
	fmt.Scanln(&letter)
	return letter
}

func main() {
	var user_input string

	fmt.Println("Helloooooo, here is a hangman game written in go lang for learing purposes")
	for {
		fmt.Println("Do you wanna start a new game?")
		fmt.Print("[Y]es [N]o ")
		fmt.Scanln(&user_input)
		if strings.ToLower(user_input) == "n" {
			fmt.Println("Bye bye!!!")
			break
		} else {
			hideWord(secret)
		}
	}
}
