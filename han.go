package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func HangmanPositions(vie *HangManData) {

	// print hangman
	f, err := os.Open("hangman.txt")
	var tab []string

	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	if vie.Attempts == 9 {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()

		fmt.Println(tab[6])

	}

	if vie.Attempts == 8 {
		fmt.Println()
		for i := 9; i <= 14; i++ {
			fmt.Println(tab[i])
		}

	}

	if vie.Attempts == 7 {
		fmt.Println()
		for i := 15; i <= 22; i++ {
			fmt.Println(tab[i])
		}
	}

	if vie.Attempts == 6 {
		fmt.Println()
		for i := 23; i <= 30; i++ {
			fmt.Println(tab[i])
		}
	}

	if vie.Attempts == 5 {
		fmt.Println()
		for i := 31; i <= 38; i++ {
			fmt.Println(tab[i])
		}
	}

	if vie.Attempts == 4 {
		fmt.Println()
		for i := 39; i <= 46; i++ {
			fmt.Println(tab[i])
		}
	}

	if vie.Attempts == 3 {
		fmt.Println()
		for i := 47; i <= 54; i++ {
			fmt.Println(tab[i])
		}
	}

	if vie.Attempts == 2 {
		fmt.Println()
		for i := 55; i <= 62; i++ {
			fmt.Println(tab[i])
		}
	}

	if vie.Attempts == 1 {
		fmt.Println()
		for i := 63; i <= 70; i++ {
			fmt.Println(tab[i])
		}
	}

	if vie.Attempts == 0 {
		fmt.Println()
		for i := 72; i <= 78; i++ {
			fmt.Println(tab[i])
		}
		fmt.Println("Vous avez perdu!")
	}

}
