// /$$                      /$$$$$$$  /$$$$$$ /$$   /$$  /$$$$$$
// | $$                     | $$__  $$|_  $$_/| $$$ | $$ /$$__  $$
// | $$  /$$$$$$   /$$$$$$  | $$  \ $$  | $$  | $$$$| $$| $$  \__/
// | $$ |____  $$ /$$__  $$ | $$$$$$$/  | $$  | $$ $$ $$| $$ /$$$$
// | $$  /$$$$$$$| $$  \__/ | $$____/   | $$  | $$  $$$$| $$|_  $$
// | $$ /$$__  $$| $$       | $$        | $$  | $$\  $$$| $$  \ $$
// | $$|  $$$$$$$| $$       | $$       /$$$$$$| $$ \  $$|  $$$$$$/
// |__/ \_______/|__/       |__/      |______/|__/  \__/ \______/

package main

import (
	"fmt"
	"os"
	"time"
)

const (
	colorBlack   = "\033[30m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"

	colorReset = "\033[0m"
)

func main() {
	check_args()
	print_logo()
	var targetIP string = get_ip()

	const starting_text string = "[ + ] 5t4rt3d h4ck1ng 0f: "

	display_text(starting_text, colorGreen)
	display_text_ln(targetIP, colorRed)
}

func check_args() {
	if len(os.Args) < 2 {
		fmt.Println("[!] 3RR0R! N0 M41NFR4M3 T0 H4X!")
		fmt.Println("USAGE: larPING <ip>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("[!] 3RR0R! 0NLY 1 M41NFR4M3 N0W :c ")
		fmt.Println("USAGE: larPING <ip>")
		os.Exit(1)
	}
}

func get_ip() string {
	return os.Args[1]
}

func print_logo() {
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

	fmt.Print(colorCyan)
	for i := 0; i < len(asciiLogo); i++ {
		fmt.Printf("%c", asciiLogo[i])
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(colorReset)
}

func display_text(text, color string) {
	fmt.Print(color)
	for i := range len(text) {
		fmt.Printf("%c", text[i])
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("%s", colorReset)
}

func display_text_ln(text string, color string) {
	display_text(text, color)
	fmt.Println()
}

// [clswhre@arch ~]$ ping 1.1.1.1
// PING 1.1.1.1 (1.1.1.1) 56(84) bytes of data.
// 64 bytes from 1.1.1.1: icmp_seq=1 ttl=59 time=7.05 ms
// 64 bytes from 1.1.1.1: icmp_seq=2 ttl=59 time=18.4 ms
// 64 bytes from 1.1.1.1: icmp_seq=3 ttl=59 time=9.83 ms
// 64 bytes from 1.1.1.1: icmp_seq=4 ttl=59 time=8.33 ms
// 64 bytes from 1.1.1.1: icmp_seq=5 ttl=59 time=23.4 ms
// 64 bytes from 1.1.1.1: icmp_seq=6 ttl=59 time=17.2 ms
// 64 bytes from 1.1.1.1: icmp_seq=7 ttl=59 time=20.9 ms
// 64 bytes from 1.1.1.1: icmp_seq=8 ttl=59 time=21.0 ms

// --- 1.1.1.1 ping statistics ---
// 8 packets transmitted, 8 received, 0% packet loss, time 7008ms
// rtt min/avg/max/mdev = 7.050/15.766/23.370/5.997 ms
