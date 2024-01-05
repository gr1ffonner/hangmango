package main

var hangmanStates = map[int]string{
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
