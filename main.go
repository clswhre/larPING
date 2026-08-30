package main

import (
	"fmt"
	"os"
	"time"
)

const colorGreen = "\033[32m"
const colorRed = "\033[31m"
const colorCyan = "\033[36m"
const colorReset = "\033[0m"

func main() {
	print_logo()
	var targetIP string = check_args()

	const starting_text string = "[ + ] Started haxing da: "
	
	fmt.Print(colorGreen)
	for i := 0; i < len(starting_text); i++ {
		fmt.Printf("%c",starting_text[i])
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("%s", colorReset)

	fmt.Print(colorRed)
	for i := 0; i < len(targetIP); i++ {
		fmt.Printf("%c",targetIP[i])
		time.Sleep(50 * time.Millisecond)
		if i == ( len(targetIP) - 1 ) { fmt.Println() }
	}
	fmt.Printf("%s", colorReset)
}


func check_args() string{
	if len(os.Args) < 2 {
		fmt.Println("[!] ERROR! NO MAINFRAME TO HAX!")
		fmt.Println("USAGE: larPING <ip>")
		os.Exit(1)
	}
	return os.Args[1]
}

func print_logo(){
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
		fmt.Printf("%c",asciiLogo[i])
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(colorReset)
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