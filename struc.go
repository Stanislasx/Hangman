package hangman

type HangManData struct {
	motAleatoire     string     // Word composed of '_', ex: H_ll_
	HideWord         string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Input            string     // Letter entered by the player
}
