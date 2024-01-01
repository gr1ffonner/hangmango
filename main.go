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

func hideWord(word string) {
	word_lenght := len(word)
	hided_word := strings.Repeat(star, word_lenght)
	fmt.Println(hided_word)
}

func unhideWord(word *string, letter string) {
	for _, v := range *word {
		if string(v) != letter {
			*word = strings.Replace(*word, string(v), star, 1)
		}
	}
	fmt.Println(*word)
}

func typeLetter() string {
	var letter string
	fmt.Println("Type letter: ")
	fmt.Scanln(&letter)
	return letter
}

func main() {
	var user_input string

	fmt.Println("Helloooooo, here is a hangman game written in golang for learing purposes")
	for {
		fmt.Println("Do you wanna start a new game?")
		fmt.Print("[Y]es [N]o ")
		fmt.Scanln(&user_input)
		if strings.ToLower(user_input) == "n" {
			fmt.Println("Bye bye!!!")
			break
		} else {
			fmt.Println("Sooooo, let's start!!!")
			hideWord(secret)
			fmt.Println(hangman)
			for {
				letter := typeLetter()
				if strings.Contains(secret, letter) {
					unhideWord(&secret, letter)
				} else {
					fmt.Println("There is not this letter in secret word")
				}
			}
		}
	}
}
