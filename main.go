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
	"os/signal"
	"time"

	probing "github.com/prometheus-community/pro-bing"
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

	print_text(starting_text, colorGreen)
	print_text_ln(targetIP, colorRed)

	time.Sleep(500 * time.Millisecond)
	ping_target(targetIP)
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
	for i := range len(asciiLogo) {
		fmt.Printf("%c", asciiLogo[i])
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(colorReset)
}

func print_text(text, color string) {
	fmt.Print(color)
	for i := range len(text) {
		fmt.Printf("%c", text[i])
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("%s", colorReset)
}

func print_text_ln(text string, color string) {
	print_text(text, color)
	fmt.Println()
}

func ping_target(ip string) {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		errMsg := fmt.Sprintf("[!] 3RR0R: %v\n", err)
		print_text_ln(errMsg, colorRed)
		return
	}

	pinger.SetPrivileged(false)

	pinger.OnRecv = func(pkt *probing.Packet) {

		// add some fun things here

		fmt.Printf("%s[DATA_RECV] %d bytes from %s: icmp_seq=%d time=%v%s\n",
			colorWhite, pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt.Round(time.Millisecond), colorReset)
	}

	// listen for ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		pinger.Stop()
	}()

	initMsg := fmt.Sprintf("[ + ] 1N1T14T1NG C0NN3CT10N T0 %s...", ip)
	print_text_ln(initMsg, colorYellow)

	err = pinger.Run()
	if err != nil {
		failMsg := fmt.Sprintf("[ ! ] C0NN3CT10N F41L3D: %v", err)
		print_text_ln(failMsg, colorRed)
		return
	}
	stats := pinger.Statistics()
	headerMsg := fmt.Sprintf("--- %s t4rg3t st4ts ---", stats.Addr)
	print_text_ln(headerMsg, colorCyan)

	statsMsg := fmt.Sprintf("%d packets transmitted, %d received, %v%% packet loss",
		stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	print_text_ln(statsMsg, colorWhite)
}
