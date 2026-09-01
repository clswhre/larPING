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
	"time"

	"github.com/likexian/whois"
	probing "github.com/prometheus-community/pro-bing"
	flag "github.com/spf13/pflag"
)

type Config struct {
	TargetIP string
	Count    int
	Whois    bool
	Size     int
}

func main() {
	config := parseArgs()

	printLogo()

	printTypewriter(startingText, colorGreen)
	printTextLn(config.TargetIP, colorRed)

	time.Sleep(500 * time.Millisecond)

	if config.Whois {
		whoisQuery(config.TargetIP)
	}
	ping_target(config)
}

func parseArgs() Config {
	var cfg Config

	flag.IntVarP(&cfg.Count, "count", "c", 0, "Numb3r 0f p4ck37s 70 s3nd (0 = infini73)")
	flag.BoolVarP(&cfg.Whois, "whois", "w", false, "P3rf0rm WH0IS l00kup 0n 74rg37")
	flag.IntVarP(&cfg.Size, "size", "s", 56, "ICMP p4ck37 p4yl04d siz3 in by73s")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "US4G3: larPING [0pt10n5] <ip>\n\n0pt10ns:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("[!] 3RR0R! M41NFR4M3 0V3RL04D!")
		flag.Usage()
		os.Exit(1)
	}

	cfg.TargetIP = args[0]
	if net.ParseIP(cfg.TargetIP) == nil {
		fmt.Println("[!] 3RR0R! N0T V4LI8 1P :c ")
		flag.Usage()
		os.Exit(1)
	}

	return cfg
}

func ping_target(cfg Config) {
	pinger, err := probing.NewPinger(cfg.TargetIP)
	if err != nil {
		errMsg := fmt.Sprintf(errorText, err)
		printTextLn(errMsg, colorRed)
		return
	}

	pinger.SetPrivileged(false)

	pinger.OnRecv = func(pkt *probing.Packet) {

		pcktMsg := fmt.Sprintf(packetText,
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

	initMsg := fmt.Sprintf(initConnection, cfg.TargetIP)
	printTextLn(initMsg, colorYellow)

	if cfg.Count > 0 {
		pinger.Count = cfg.Count
	}

	pinger.Size = cfg.Size

	err = pinger.Run()
	if err != nil {
		failMsg := fmt.Sprintf(failText, err)
		printTextLn(failMsg, colorRed)
		return
	}
	stats := pinger.Statistics()
	headerMsg := fmt.Sprintf(headerText, stats.Addr)
	printTextLn(headerMsg, colorCyan)

	statsMsg := fmt.Sprintf(statText,
		stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	printTextLn(statsMsg, colorWhite)
	fmt.Println("[+] M1SS10N C0MPL3T3. 3X1T...")
}

func whoisQuery(domain string) {
	printTypewriter(whoisText, colorYellow)
	time.Sleep(500 * time.Millisecond)
	result, err := whois.Whois(domain)
	if err == nil {
		printInstant(result, colorBlue)
	} else {
		errMsg := fmt.Sprintf(whoisErrText, err)
		printTextLn(errMsg, colorRed)
	}
}
