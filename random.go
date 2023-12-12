package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func ChooseRandomWord(filename string) (string, error) {

	// choose random word from file

	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Impossible d'ouvrir le fichier %s : %v", filename, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	words := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		wordList := strings.Fields(line)
		words = append(words, wordList...)
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Erreur lors de la lecture du fichier %s : %v", filename, err)
	}
	if len(words) == 0 {
		return "", fmt.Errorf("Le fichier %s est vide", filename)
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	return words[randomIndex], nil
}
