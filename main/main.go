package main

import "github.com/Stanislasx/hangman"

func main() {
	// Créez une instance de HangManData pour fournir à la fonction Test
	vie := hangman.HangManData{}
	hangman.Test(&vie)
}
