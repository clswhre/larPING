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
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func main() {
	parseArgs()
	printLogo()
	var targetIP string = parseArgs()

	const startingText string = "[ + ] 5t4rt3d h4ck1ng | ip: "

	printTypewriter(startingText, colorGreen)
	printTextLn(targetIP, colorRed)

	time.Sleep(500 * time.Millisecond)
	ping_target(targetIP)
}

func parseArgs() string {
	// no ip
	if len(os.Args) < 2 {
		fmt.Println("[!] 3RR0R! N0 M41NFR4M3 T0 H4X!")
		fmt.Println("USAGE: larPING <ip>")
		os.Exit(1)
	}

	// ip + some other things(?)
	if len(os.Args) > 2 {
		fmt.Println("[!] 3RR0R! 0NLY 1 M41NFR4M3 N0W :c ")
		fmt.Println("USAGE: larPING <ip>")
		os.Exit(1)
	}

	// actually parse
	if net.ParseIP(os.Args[1]) == nil {
		fmt.Println("[!] 3RR0R! N0T V4LI8 1P :c ")
		fmt.Println("USAGE: larPING <ip>")
		os.Exit(1)
	}
	return os.Args[1]
}

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

func ping_target(ip string) {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		errMsg := fmt.Sprintf("[!] 3RR0R: %v\n", err)
		printTextLn(errMsg, colorRed)
		return
	}

	pinger.SetPrivileged(false)

	pinger.OnRecv = func(pkt *probing.Packet) {

		pcktMsg := fmt.Sprintf(">>> %d 8yt35 fr0m %s | 1cmp_53q=%d | 71m3=%v",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt.Round(time.Millisecond))
		printInstant(pcktMsg, colorBlue)

		// add some fun things here
		printInstant(phrases[rand.IntN(len(phrases))], colorYellow)

		if rand.IntN(4) == 0 {
			printLarpLogos()
		}
		printInstant("----------------------------------------", colorCyan)
	}

	// listen for ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		pinger.Stop()
	}()

	initMsg := fmt.Sprintf("[ + ] 1N1T14T1NG C0NN3CT10N T0 %s...", ip)
	printTextLn(initMsg, colorYellow)

	err = pinger.Run()
	if err != nil {
		failMsg := fmt.Sprintf("[ ! ] C0NN3CT10N F41L3D: %v", err)
		printTextLn(failMsg, colorRed)
		return
	}
	stats := pinger.Statistics()
	headerMsg := fmt.Sprintf("\n--- %s t4rg3t st4ts ---", stats.Addr)
	printTextLn(headerMsg, colorCyan)

	statsMsg := fmt.Sprintf("%d packets transmitted, %d received, %3.2f%% packet loss",
		stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	printTextLn(statsMsg, colorWhite)
	fmt.Println("[+] M1SS10N C0MPL3T3. 3X1T...")
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
