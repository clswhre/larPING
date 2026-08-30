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
	"math/rand/v2"
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

var phrases = []string{
	"[*] 8YP4551NG M41NFR4M3 F1R3W4LL...",
	"[!] 1NJ3CT1NG P4YL04D...",
	"[*] D3CRYPT1NG P455W0RD H45H35...",
	"[+] D0WN104D1NG S3CR3T D4T4...",
	"[!] TR4C3 D3T3CT3D, R0UT1NG...",
	"[*] 1N1T14L1Z1NG 3NCRYPT10N BYP455...",
	"[!] 3XPL01T1NG 0V3RFL0W VULN...",
	"[*] 3NUM3R4T1NG H1DD3N D1R3CT0R13S...",
	"[+] 4CC3SS GR4NT3D. W3LC0M3 T0 TH3 M41NFR4M3...",
	"[!] 5Y5T3M 4L3RT: 1NTRUD3R D3T3CT3D. 3V4D1NG...",
	"[*] 1NJ3CT1NG R3V3R5E 5H3LL C0D3...",
	"[+] P4YLOAD D3L1V3R3D. 3X3CUT1NG...",
	"[!] F1R3W4LL RUL3S BYP4553D...",
	"[*] CR4CK1NG W1F1 K3Y 4LG0R1THM...",
	"[+] D4T4B4S3 DUMP 1N PR0GR3SS...",
	"[!] R00TK1T 1NST4LL3D. P3R515T3NC3 3N4BL3D...",
	"[*] 5C4NN1NG P0RT5 1337-65535...",
	"[+] 5H3LL 4CC3SS 0BT41N3D. 3SC4L4T1NG PR1V1L3G3S...",
	"[!] 1D5 3V4D3D. CH4NG1NG M4C 4DDR3SS...",
	"[*] 3R4S1NG L0G5. N0 TR4C3 L3FT...",
	"[+] M1SS10N C0MPL3T3. 3X1T...",
}

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
		print_larp_text()
		pcktMsg := fmt.Sprintf("%d 8yt35 fr0m %s | 1cmp_53q=%d | 71m3=%v",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt.Round(time.Millisecond))
		print_text_ln(pcktMsg, colorBlue)
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

	statsMsg := fmt.Sprintf("%d packets transmitted, %d received, %3.2f%% packet loss",
		stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	print_text_ln(statsMsg, colorWhite)
}

func print_larp_text() {

	randomInt := rand.IntN(len(phrases))
	fmt.Println("\n", phrases[randomInt])
}
