package main

import (
	"fmt"
	"hangman"
	"math/rand"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Veuillez fournir un nom de fichier de mots.")
		return
	}
	// List of file names containing words
	motAleatoire, err := hangman.ChooseRandomWord(os.Args[1])
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Println("Mot aléatoire :")
	// Reveal n random letters in word
	n := len(motAleatoire)/2 - 1
	revealed := make([]bool, len(motAleatoire))
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(motAleatoire))
		if !revealed[randomIndex] {
			revealed[randomIndex] = true
		} else {
			i--
		}
	}
	// Show word with revealed letters
	for i, c := range motAleatoire {
		if revealed[i] {
			fmt.Printf("%c ", c)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
	var vie hangman.HangManData
	vie.Attempts = 10
	for vie.Attempts > 0 {
		// Read a letter from the player
		letter := hangman.ReadLetter()
		found := false
		for i, value := range motAleatoire {
			if string(value) == letter {
				revealed[i] = true
				found = true
			}
		}

		if letter == motAleatoire {
			fmt.Println("Vous avez gagné!")
			break
		} else if len(letter) > 1 {
			vie.Attempts -= 2
		}
		if letter != motAleatoire && found == false {
			vie.Attempts--
			fmt.Printf("La lettre %s n'est pas dans le mot. Il vous reste %d essais.\n", letter, vie.Attempts)
		}
		if vie.Attempts == 0 {
			fmt.Println("Vous avez perdu!")
			fmt.Println("Le mot était :", motAleatoire)
			break
		}

		// Show word with revealed letters
		for i, c := range motAleatoire {
			if revealed[i] {
				fmt.Printf("%c ", c)
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
		// Check if the game is over
		if hangman.HideWord(motAleatoire, revealed) == motAleatoire {
			fmt.Println("Vous avez gagné!")
			fmt.Println("Vous avez trouvé le mot", motAleatoire)
			break
		}
		hangman.HangmanPositions(&vie)
	}
	fmt.Println("mot aléatoire : ", motAleatoire)
}
