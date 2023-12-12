package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Test(vie *HangManData) {
	if len(os.Args) != 2 {
		fmt.Println("Veuillez fournir un nom de fichier de mots.")
		return
	}

	// Liste des noms de fichiers contenant des mots
	motAleatoire, err := ChooseRandomWord(os.Args[1])
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Println("Mot aléatoire :")

	// Révéler n lettres aléatoires dans le mot
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

	// Afficher le mot avec les lettres révélées
	for i, c := range motAleatoire {
		if revealed[i] {
			fmt.Printf("%c ", c)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()

	vie.Attempts = 10
	for vie.Attempts > 0 {
		// Lire une lettre du joueur
		letter := ReadLetter()
		found := false

		for i, value := range motAleatoire {
			if string(value) == letter {
				revealed[i] = true
				found = true
			}
		}

		if string(motAleatoire) == letter {
			fmt.Println("Vous avez gagné!")
			break
		} else if len(letter) > 1 {
			vie.Attempts -= 2
		}

		if string(motAleatoire) != letter && !found {
			vie.Attempts--
			fmt.Printf("La lettre %s n'est pas dans le mot. Il vous reste %d essais.\n", letter, vie.Attempts)
		}

		if vie.Attempts == 0 {
			fmt.Println("Vous avez perdu!")
			fmt.Println("Le mot était :", motAleatoire)
			break
		}

		// Afficher le mot avec les lettres révélées
		for i, c := range motAleatoire {
			if revealed[i] {
				fmt.Printf("%c ", c)
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()

		// Vérifier si le jeu est terminé
		if HideWord(motAleatoire, revealed) == motAleatoire {
			fmt.Println("Vous avez gagné!")
			fmt.Println("Vous avez trouvé le mot", motAleatoire)
			break
		}

		HangmanPositions(vie)
	}

	fmt.Println("Mot aléatoire : ", motAleatoire)
}
