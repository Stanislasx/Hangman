package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLetter() string {

	// read letter from player

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrer une lettre : ")
	letter, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur:", err)
		return ""
	}
	return strings.TrimSpace(letter)
}
