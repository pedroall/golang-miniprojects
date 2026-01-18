package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Print("Type a phrase and a random word will be chosen: \n")

	// Configurating the Reader

	reader := bufio.NewReader(os.Stdin)

	// Reads until finds a breakline

	text, texterror := reader.ReadString('\n')

	if texterror != nil {
		fmt.Println("Error: ", texterror)
	} else {
		text = strings.TrimSpace(text)
		words := strings.Fields(text) // Could've used strings.Split

		randomNumber := rand.Intn(len(words))
		randomWord := words[randomNumber]
		fmt.Println(randomWord)
	}
}
