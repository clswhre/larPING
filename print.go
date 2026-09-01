package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

func printLogo() {
	const asciiLogo string = `
	/$$                      /$$$$$$$  /$$$$$$ /$$   /$$  /$$$$$$ 
	| $$                     | $$__  $$|_  $$_/| $$$ | $$ /$$__  $$
	| $$  /$$$$$$   /$$$$$$  | $$  \ $$  | $$  | $$$$| $$| $$  \__/
	| $$ |____  $$ /$$__  $$ | $$$$$$$/  | $$  | $$ $$ $$| $$ /$$$$
	| $$  /$$$$$$$| $$  \__/ | $$____/   | $$  | $$  $$$$| $$|_  $$
	| $$ /$$__  $$| $$       | $$        | $$  | $$\  $$$| $$  \ $$
	| $$|  $$$$$$$| $$       | $$       /$$$$$$| $$ \  $$|  $$$$$$/
	|__/ \_______/|__/       |__/      |______/|__/  \__/ \______/ 
	`

	fmt.Print(colorBlue)
	for i := range len(asciiLogo) {
		fmt.Printf("%c", asciiLogo[i])
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(colorReset)
}

func printTypewriter(text, color string) {
	fmt.Print(color)
	for i := range len(text) {
		fmt.Printf("%c", text[i])
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("%s", colorReset)
}

func printTextLn(text string, color string) {
	printTypewriter(text, color)
	fmt.Println()
}

func printInstant(text, color string) {
	fmt.Printf("%s%s%s\n", color, text, colorReset)
}

func printLarpText() {

	randomInt := rand.IntN(len(phrases))
	fmt.Println("\n", phrases[randomInt])
}

func printLarpLogos() {
	asciiLogo := logos[rand.IntN(len(logos))]
	asciiLogo = strings.ReplaceAll(asciiLogo, "@", "`")

	randomColor := logoColors[rand.IntN(len(logoColors))]

	fmt.Print(randomColor)
	fmt.Println(asciiLogo)
	fmt.Print(colorReset)
}
