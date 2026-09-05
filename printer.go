package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

func printColor(color, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("%s%s%s\n", color, msg, colorReset)
}

func printTypewriter(delay time.Duration, color, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Print(color)
	for _, r := range msg {
		fmt.Printf("%c", r)
		if delay > 0 {
			time.Sleep(delay)
		}
	}
	fmt.Print(colorReset)
}

func printLarpLogos() {
	logo := strings.ReplaceAll(logos[rand.IntN(len(logos))], "@", "`")
	color := logoColors[rand.IntN(len(logoColors))]
	printColor(color, "%s", logo)
}
